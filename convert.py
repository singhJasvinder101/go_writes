import cv2 as cv
import numpy as np
import os
from glob import glob

def process_image(image_path):
    img = cv.imread(image_path, cv.IMREAD_GRAYSCALE)
    if img is None:
        print(f"Error: Could not read {image_path}")
        return

    _, thresh = cv.threshold(img, 197, 255, cv.THRESH_BINARY)

    scale_factor = 3
    height, width = thresh.shape[:2]
    new_size = (int(width * scale_factor), int(height * scale_factor))

    resized_img = cv.resize(thresh, new_size, interpolation=cv.INTER_NEAREST)

    padding_top = 30
    padding_bottom = 30
    padded_img = cv.copyMakeBorder(resized_img, padding_top, padding_bottom, 0, 0, cv.BORDER_CONSTANT, value=255)

    cv.imwrite(image_path, padded_img)
    print(f"Processed and saved: {image_path}")

def process_all_images(root_folder):
    image_files = glob(os.path.join(root_folder, "**", "*.png"), recursive=True)

    if not image_files:
        print("No PNG images found in the specified directory.")
        return

    print(f"Found {len(image_files)} images. Processing...")

    for image_file in image_files:
        process_image(image_file)

    print("Processing complete!")

train_copy_folder = "archive/train copy"

process_all_images(train_copy_folder)
