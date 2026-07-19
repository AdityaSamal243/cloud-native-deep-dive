import json
import os

INPUT_FILE = "ApplicationList.json"
OUTPUT_FILE = "Updated_ApplicationList.json"

REQUIRED_ENTRIES = [
    { "application_id": "60", "text": "Adhoc/Widget/Spanning Tree" },
    { "application_id": "61", "text": "Adhoc/Widget/Prediction Line" },
    { "application_id": "62", "text": "Package/ExportPdf" },
    { "application_id": "63", "text": "Adhoc/Add Formula Button" },
    { "application_id": "64", "text": "Adhoc/Cog Search" },
    { "application_id": "65", "text": "Administrator/Modules" },
    { "application_id": "66", "text": "Administrator/Create Layout" },
    { "application_id": "67", "text": "Dashboard Viewer" },
    { "application_id": "68", "text": "Browse Dashboard" },
    { "application_id": "69", "text": "New Dashboard Designer" },
    { "application_id": "70", "text": "AI Playground" },
    { "application_id": "71", "text": "Microservice" },
    { "application_id": "72", "text": "Template Designer" },
    { "application_id": "Designer Admin", "text": "Designer Admin" },
    { "application_id": "AI Dashboard", "text": "AI Dashboard" },
    { "application_id": "AI Dashboard/Dev_Controls", "text": "AI Dashboard/Dev_Controls" },
    { "application_id": "CustomTemplates", "text": "CustomTemplates" },
    { "application_id": "Data Source/Custom Truncate", "text": "Data Source/Custom Truncate" },
    { "application_id": "chat_Box", "text": "chat_Box" }
]

def generate_updated_list():
    # 1. Verify if the source file exists
    if not os.path.exists(INPUT_FILE) or os.path.getsize(INPUT_FILE) == 0:
        print(f"Error: Source file '{INPUT_FILE}' not found or empty. Generating '{OUTPUT_FILE}' with only the required entries.")
        with open(OUTPUT_FILE, 'w', encoding='utf-8') as f:
            json.dump(REQUIRED_ENTRIES, f, indent=4, ensure_ascii=False)
        return

    # 2. Parse the source data
    try:
        with open(INPUT_FILE, 'r', encoding='utf-8') as f:
            root_data = json.load(f)
    except json.JSONDecodeError as e:
        print(f"Error parsing JSON from {INPUT_FILE}: {e}")
        return

    # 3. Locate the list array dynamically
    target_list = None
    if isinstance(root_data, list):
        target_list = root_data
    elif isinstance(root_data, dict):
        for key, value in root_data.items():
            if isinstance(value, list):
                target_list = value
                print(f"Detected application array inside JSON key: '{key}'")
                break
        
        if target_list is None:
            print(f"Error: '{INPUT_FILE}' is an object, but no nested list/array was found.")
            return
    else:
        print(f"Error: Unexpected JSON root structure type: {type(root_data).__name__}")
        return

    # 4. Extract existing IDs to avoid duplicate insertions
    existing_ids = {str(item.get("application_id")) for item in target_list if isinstance(item, dict) and "application_id" in item}

    # 5. Append missing elements into the memory copy
    missing_entries_added = 0
    for entry in REQUIRED_ENTRIES:
        if str(entry["application_id"]) not in existing_ids:
            target_list.append(entry)
            print(f"Staged missing entry: {entry['application_id']} -> {entry['text']}")
            missing_entries_added += 1

    # 6. Dump everything into the new separate file
    with open(OUTPUT_FILE, 'w', encoding='utf-8') as f:
        json.dump(root_data, f, indent=4, ensure_ascii=False)
    
    print(f"\nSuccess! Original '{INPUT_FILE}' remains untouched.")
    print(f"Created/updated: '{OUTPUT_FILE}' (Added {missing_entries_added} missing entries).")

if __name__ == "__main__":
    generate_updated_list()