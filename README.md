## Gallery fullstack web application

<img src="https://github.com/Rryowa/galleram/assets/80339180/cee48a81-a6df-44e8-b1ff-9728c44975cd"/>

## API based on: #
* router: &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Chi
* database: Postgres
* db driver: &nbsp;pgx
* CSS: &emsp;&emsp;Tailwind

## Model-View-Controller (MVC) structure:
* Models - data, logic rules
* Views - render data (html/generate JSON)
* Controllers - connects all

## How I protected passwords:
* HTTPs
* Hash - HMAC(sha256) [digital signing]
* Salt - (bcrypt) [passwords]
* Time-constant functions during auth
