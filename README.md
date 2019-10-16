# uinput-event-api

Send input events

## Endpoints

### `POST /keypress`

Send keypress event

#### Request body:

```json
{ "key": "UP" }
```

#### Example with curl

`curl -X POST -H "Content-Type: application/json" -d '{"key":"UP"}' http://localhost:8085/keypress`

## Available keys

* `ESC`
* `A`
* `Z`
* `HOME`
* `UP`
* `LEFT`
* `RIGHT`
* `DOWN`

## License

Licensed under MIT ([LICENSE](LICENSE) or https://opensource.org/licenses/MIT).

## Contribution

Unless you explicitly state otherwise, any contribution intentionally
submitted for inclusion in the work by you, as defined in the MIT
license, shall be licensed as above, without any additional terms or
conditions.