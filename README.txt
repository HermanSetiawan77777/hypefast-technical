Name : Shorten Link (POST)
URL: http://localhost:8080/shorten
Request : "URL" string
Response : "URL" (shorten) 


Name : Redirect Original Link (GET)
URL: http://localhost:8080/{id}
Request : "id" string
Response : "Original Link"


Name : cek status (GET)
URL: http://localhost:8080/{id}/stats
Request : "id" string
Response : "created_at", "redirect_count" 