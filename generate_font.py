import numpy as np
import cv2
import imutils
import os

output_folder = "cropped_chars"
os.makedirs(output_folder, exist_ok=True)

image = cv2.imread("image.png")
if image is None:
    print("Error: Could not load image. Check the file path.")
    exit()

gray = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)

ret, thresh1 = cv2.threshold(gray, 127, 255, cv2.THRESH_BINARY_INV)

dilate = cv2.dilate(thresh1, None, iterations=2)
dilate = np.uint8(dilate) 
cnts = cv2.findContours(dilate.copy(), cv2.RETR_EXTERNAL, cv2.CHAIN_APPROX_SIMPLE)
cnts = imutils.grab_contours(cnts)

if not cnts:
    print("No contours found!")
    exit()

orig = image.copy()

TARGET_HEIGHT = 120
TARGET_WIDTH = 48

i = 0
for cnt in cnts:
    if cv2.contourArea(cnt) < 100:
        continue

    x, y, w, h = cv2.boundingRect(cnt)

    roi = image[y:y+h, x:x+w]
    aspect_ratio = w / h
    if aspect_ratio > (TARGET_WIDTH / TARGET_HEIGHT):  # Too wide, limit width
        new_w = TARGET_WIDTH
        new_h = int(TARGET_WIDTH / aspect_ratio)
    else: 
        new_h = TARGET_HEIGHT
        new_w = int(TARGET_HEIGHT * aspect_ratio)

    resized_roi = cv2.resize(roi, (new_w, new_h), interpolation=cv2.INTER_AREA)

    final_image = np.ones((TARGET_HEIGHT, TARGET_WIDTH, 3), dtype=np.uint8) * 255

    x_offset = (TARGET_WIDTH - new_w) // 2
    y_offset = (TARGET_HEIGHT - new_h) // 2

    final_image[y_offset:y_offset+new_h, x_offset:x_offset+new_w] = resized_roi

    cv2.imwrite(os.path.join(output_folder, f"roi_{i}.png"), final_image)

    cv2.rectangle(orig, (x, y), (x + w, y + h), (0, 255, 0), 2)

    i += 1

print(f"Segmented characters saved in '{output_folder}' folder.")

cv2.imshow("Segmented Characters", orig)
cv2.waitKey(0)
cv2.destroyAllWindows()
