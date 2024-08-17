# ReqMatch
golang HTTP service to match http request, that was received in 10 seconds interval

## Testing
Run 1000 request, with 100 parallel jobs

`
seq 1 1000 | xargs -Iname  -P100  curl -X POST "http://localhost:3000/join"
curl "http://localhost:3000/stats" | jq
`