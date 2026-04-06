# Order Calculator


### Synopsis

This repo contains take home assignment for Software Engineering role.

Given an array of available package sizes that cannot be split, this application uses Dynamic Programming technique to calculate such an order configuration, so that for any order size two things are ensured:
1) Minimal amount of items is shipped over required order size
2) Minimal amount of packs will be loaded into order

> Example 1:  
&nbsp;&nbsp;&nbsp;&nbsp;Available Packs: [25, 50, 100, 200]  
&nbsp;&nbsp;&nbsp;&nbsp;Order Size: 174  
&nbsp;&nbsp;&nbsp;&nbsp;Result: {100: 1, 50: 1, 25: 1}  
Explanation:  
&nbsp;&nbsp;&nbsp;&nbsp;Overshipped items = 100x1 + 50x1 + 25x1 - 174 = 175 - 174 = 1  
&nbsp;&nbsp;&nbsp;&nbsp;Packs sent = 3

> Example 2:  
&nbsp;&nbsp;&nbsp;&nbsp;Available Packs: [25, 50, 100, 200]  
&nbsp;&nbsp;&nbsp;&nbsp;Order Size: 176 
&nbsp;&nbsp;&nbsp;&nbsp;Result: {200: 1}  
Explanation:  
&nbsp;&nbsp;&nbsp;&nbsp;Overshipped items = 200 174 = 26  
&nbsp;&nbsp;&nbsp;&nbsp;Packs sent = 1  
This is better in terms of packs sent than sending {100: 2} or {500: 4} or {250:8}, or any other permutation of pack sizes under 200 due to same overshipping amount, but smaller packaging overhead.

### Demo
[demo.webm](https://github.com/user-attachments/assets/28fc31d8-13b9-479a-8225-65c3c6576dc4)


### Usage

The app is accessible either via API calls or UI on [Render](https://render.com/). Please note,that due to free tier it make take a while to spin-up the app.

For UI access:  
[https://order-calculator-1z6f.onrender.com/](https://order-calculator-1z6f.onrender.com/)

For API:
* Get available package sizes
```sh
  curl https://order-calculator-1z6f.onrender.com/api/packs
```

* Calculate order configuration for order of size `order_size`
```sh
    curl -X POST "https://order-calculator-1z6f.onrender.com/api/order" -H "Content-Type: application/json" -d '{"order_size": <order_size>}' 

```

### Development
The app is containerized, and therefore can be run locally with the following command:
```
docker compose up --watch
```

The app is also configurable via `config.json` file. Where you can set either `port` on which app will run or `pack_sizes` that will be used to determine optimal order configuration.
