from fastapi import FastAPI
from fastapi.responses import JSONResponse
from pydantic import BaseModel

app = FastAPI(
    title="Carmasearch Scraper API",
    version="1.0.0",
    description="Carmasearch Scraper API"
)

@app.get("/")
async def root():
    res:dict = {
        "status": "Running..."
    }
    return JSONResponse(content=res, status_code=200)

class RequestBody(BaseModel):
    url: str

@app.post("/")
async def scrape(req: RequestBody):
    print(req)
    return req
