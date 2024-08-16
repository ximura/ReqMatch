# ReqMatch
golang HTTP service to match http request, that was received in 10 seconds interval

## Testing
Run 200 request, with 10 parallel jobs

`
seq 1 200 | xargs -Iname  -P10  curl -X POST "http://localhost:3000/join"
`

