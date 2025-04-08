import csv
import sys

def extract_shoot_attributes(file_path):
    """Extracts and deduplicates 'shoot' attributes from a CSV file."""
    shoot_attributes = set()
    try:
        with open(file_path, mode='r', newline='', encoding='utf-8') as csvfile:
            reader = csv.DictReader(csvfile)
            for row in reader:
                if 'shoot' in row:
                    shoot_attributes.add(row['shoot'])
    except FileNotFoundError:
        print(f"File not found: {file_path}")
    except Exception as e:
        print(f"An error occurred while processing {file_path}: {e}")
    return list(shoot_attributes)

if __name__ == "__main__":
    if len(sys.argv) != 4:
        print("Usage: python csv_formatter.py <expired_certs_csv> <about_to_expire_certs_csv> <manual_installations_csv>")
        sys.exit(1)

    # Paths to your CSV files from execution arguments
    expired_certs_csv = sys.argv[1]
    about_to_expire_certs_csv = sys.argv[2]
    manual_installations_csv = sys.argv[3]

    # Extract and deduplicate 'shoot' attributes
    expired_certs_shoots = extract_shoot_attributes(expired_certs_csv)
    about_to_expire_shoots = extract_shoot_attributes(about_to_expire_certs_csv)
    manual_installations_shoots = extract_shoot_attributes(manual_installations_csv)

    # Output the results with newline characters
    print("Expired Certs Shoots:")
    print("\n".join(expired_certs_shoots))
    print("\nAbout To Expire Certs Shoots:")
    print("\n".join(about_to_expire_shoots))
    print("\nManual Installations Shoots:")
    print("\n".join(manual_installations_shoots))