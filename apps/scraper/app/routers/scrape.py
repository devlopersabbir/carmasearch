from fastapi import APIRouter, HTTPException

router = APIRouter(
    prefix="/scrape",
    tags=["Scraping"],
)

@router.post("/", response_model=ScrapeResponse)
async def scrape_website(payload: ScrapeRequest):
    try:
        result = await run_in_thread(
            ScraperService.scrape,
            payload.url,
        )
        return result

    except Exception as e:
        raise HTTPException(
            status_code=400,
            detail=str(e),
        )