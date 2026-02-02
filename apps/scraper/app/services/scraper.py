import httpx
from bs4 import BeautifulSoup
from core.config import DEFAULT_HEADERS, REQUEST_TIMEOUT


class ScraperService:
    @staticmethod
    def fetch_html(url: str) -> str:
        with httpx.Client(
            headers=DEFAULT_HEADERS,
            timeout=REQUEST_TIMEOUT,
            follow_redirects=True,
        ) as client:
            response = client.get(url)
            response.raise_for_status()
            return response.text

    @staticmethod
    def parse_title(html: str) -> str:
        soup = BeautifulSoup(html, "lxml")
        title = soup.find("title")
        return title.text.strip() if title else "No title found"

    @classmethod
    def scrape(cls, url: str) -> dict:
        html = cls.fetch_html(url)
        title = cls.parse_title(html)

        return {
            "url": url,
            "title": title,
        }
