#!/usr/bin/env python3

import sys
import requests
from bs4 import BeautifulSoup

def fetch_and_extract(url):
    headers = {
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3",
        "Accept-Language": "en-US,en;q=0.9",
        "Accept-Encoding": "gzip, deflate, br",
        "Connection": "keep-alive",
        "Upgrade-Insecure-Requests": "1"
    }

    response = requests.get(url, headers=headers)
    response.raise_for_status()

    soup = BeautifulSoup(response.text, 'html.parser')

    # Custom function to filter div elements by class name
    def has_blogCard_card_class(tag):
        return tag.name == 'div' and tag.has_attr('class') and any('blogCard_card' in cls for cls in tag['class'])

    # Find all div elements with class names containing 'blogCard_card'
    blog_cards = soup.find_all(has_blogCard_card_class)

    # Initialize list to store extracted data
    extracted_data = []

    # Extract data from each blog card div
    for card in blog_cards:
        # Initialize dictionary for each card
        card_data = {}

        # Extract title
        title_anchor = card.find('a', class_=lambda c: c and 'blogCard_blog_title' in c)
        if title_anchor:
            card_data['Title'] = title_anchor.get_text()

        # Extract image URL
        image_tag = card.find('img', class_=lambda c: c and 'blogCard_image' in c)
        if image_tag and image_tag.has_attr('src'):
            card_data['Image'] = image_tag['src']

        # Extract URL
        url_anchor = card.find('a', class_=lambda c: c and 'blogCard_blog_title' in c)
        if url_anchor and url_anchor.has_attr('href'):
            card_data['URL'] = url_anchor['href']

        # Add card data to the list
        extracted_data.append(card_data)

    return extracted_data


def main():
    if len(sys.argv) != 2:
        print("Usage: fetch_and_extract.py <URL>")
        sys.exit(1)

    url = sys.argv[1]
    extracted_data = fetch_and_extract(url)

    # Print extracted_data in a formatted way
    for item in extracted_data:
        print("Title:", item.get('Title', 'N/A'))
        print("Image:", item.get('Image', 'N/A'))
        print("URL:", item.get('URL', 'N/A'))
        print()  # Add a new line between each dictionary item

if __name__ == "__main__":
    main()

