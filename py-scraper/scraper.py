import requests
from bs4 import BeautifulSoup
import json

pageIdx = 1
url = 'https://books.toscrape.com/catalogue/page-' + str(pageIdx) + '.html'
priceIdx = -1

# create .json file
file_path = "./products.json"
data = {}
data['data'] = []

while pageIdx < 11:
    resp = requests.get(url)
    soup = BeautifulSoup(resp.content, 'html.parser')

    # fetch data from target site
    result = soup.find_all('img', attrs={'class':'thumbnail'})
    price = soup.find_all('p', attrs={'class':'price_color'})

    # insert data into .json file
    for i in result:
        priceIdx += 1
        data['data'].append({
            "title": i.attrs['alt'],
            "url": "https://books.toscrape.com/" + i.attrs['src'],
            "price": price[priceIdx].text,
        })
    priceIdx = -1
    pageIdx += 1

with open(file_path, 'w') as outfile:
    json.dump(data, outfile, indent=4)

# task took avg 18s
