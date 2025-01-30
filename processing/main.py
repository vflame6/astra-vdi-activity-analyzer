import torch
from PIL import Image, ImageDraw

# Load YOLOv5 model
model = torch.hub.load("ultralytics/yolov5", "yolov5x6")

image_path = "screenshot.png"
image = Image.open(image_path)

# Perform inference on the image
results = model(image_path)
results.show()

# Extract detection results
detections = results.xyxy[0]  # Get the detected bounding boxes

# Prepare to draw bounding boxes
image_draw = image.copy()
draw = ImageDraw.Draw(image_draw)

print("Identified Windows:")

# Loop through each detection
for det in detections:
    x1, y1, x2, y2, conf, cls = det.tolist()
    # Convert coordinates to integers
    x1, y1, x2, y2 = int(x1), int(y1), int(x2), int(y2)
    print(f"Window detected at: ({x1}, {y1}), ({x2}, {y2})")
    # Draw a rectangle on the image
    draw.rectangle([x1, y1, x2, y2], outline="red", width=3)

# Display the image with bounding boxes
image_draw.show()
