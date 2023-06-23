from fastapi import FastAPI, Response, Request, HTTPException
from fastapi.staticfiles import StaticFiles
from fastapi.middleware.cors import CORSMiddleware
import os

app = FastAPI()

app.mount("/static", StaticFiles(directory="static"), name="static")

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],  # Allows all origins
    allow_credentials=True,
    allow_methods=["*"],  # Allows all methods
    allow_headers=["*"],  # Allows all headers
)

@app.get("/")
def read_root():
    return {"Hello": "World"}

@app.get("/files/{file_path:path}")
async def read_file(request: Request, file_path: str):
    file_location = f"static/{file_path}"
    if not os.path.isfile(file_location):
        raise HTTPException(status_code=404, detail="File not found")

    range = request.headers.get('range')
    start, end = range.split("=")[1].split("-")
    start, end = int(start), int(end)

    print(f"Client requests byte range: {start}-{end}")
    data = None
    with open(file_location, "rb") as f:
        f.seek(start)
        data = f.read((end + 1) - start)

    if data is None:
        raise HTTPException(status_code=416, detail="Range not satisfiable")

    response = Response(content=data, media_type="video/mp4", status_code=206)
    response.headers["Content-Range"] = f"bytes {start}-{end or ''}/{os.path.getsize(file_location)}"
    response.headers["Accept-Ranges"] = "bytes"
    return response
