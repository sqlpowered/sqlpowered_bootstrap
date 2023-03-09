Goals
===========

make simplifying assumptions
Simplify syntax for data fetching use-case


```json
{
    "select" :[
        ["customers.name"], 
        ["customers.country"]
    ],
    "from": ["customers"],
    "left_join": [
        ["orders.customer_id", "jsql.eq", "customers.id"]
    ],
    "where": [
        ["customers.name", "jsql.neq", "Johnny droptables"]
    ],
    "order_by" : ["customers.name", "customers.country", "jsql.asc"]
}
```
Prefixing operators and functions makes them easier to identify

Functions, let's have a crack!
```json
{
    "select" :[
        ["customers.name"],
        ["jsql.max(", "order.order_amount", ")"]
    ],
    "from": ["customers"],
    "left_join": [
        ["orders.customer_id", "jsql.eq", "customers.id"]
    ],
    "where": [
        ["customers.name", "jsql.neq", "Johnny droptables"]
    ],
    "order_by" : ["customers.name", "customers.country", "jsql.asc"]
}
```

Let's try something more complex

```sql
SELECT
	product_id,
	products.name,
	(SUM(sales.units) * (products.price - products.cost)) AS profit
FROM
	products
	LEFT JOIN sales USING (product_id)
WHERE
	sales.date > CURRENT_DATE - INTERVAL '4 weeks'
GROUP BY
	product_id,
	products.name,
	products.price,
	products.cost
HAVING
	SUM(products.price * sales.units) > 5000;

```


```json
{
    "select" :[
        ["product_id"],
        ["products.name"],
        ["(", "jsql.mult(", "jsql.sum(", "sales.units", ")", 
         "jsql.sub(", "products.price", "products.cost", ")",")", ")", "jsql.as(", "profit", ")"]
    ],
    "from": ["products"],
    "left_join": [
        ["sales", "jsql.using", "product_id"]
    ],
    "where": [
        ["sales.date", "jsql.gt", "jsql.sub(", "CURRENT_DATE", "INTERVAL '4 weeks'", ")"]
    ],
    "group_by" : [
        ["product_id"],
        ["products.name"],
        ["products.price"],
        ["products.cost"]
    ],
    "having" :[
        ["jsql.sum(", "jsql.mult(", "products.price", "sales.units", ")", ")", 
        "jsql.gt(", "5000", ")"]
    ]
}
```
This uses functions to make the sql easier to parse and is already tokenised
tokenised is quite ugly let's try taking that on:


```json
{
    "select" :[
        ["product_id"],
        ["products.name"],
        ["( jsql.mult( jsql.sum(sales.units), jsql.sub( products.price, products.cost ) ) ) jsql.as(profit)"]
    ],
    "from": ["products"],
    "left_join": [
        ["sales jsql.using product_id"]
    ],
    "where": [
        ["sales.date jsql.gt jsql.sub( CURRENT_DATE, INTERVAL '4 weeks' )"]
    ],
    "group_by" : [
        ["product_id"],
        ["products.name"],
        ["products.price"],
        ["products.cost"]
    ],
    "having" :[
        ["jsql.sum( jsql.mult( products.price, sales.units ) ) jsql.gt(5000)"]
    ]
}
```


```json
{
    "select" :[
        "product_id",
        "products.name",
        "(jsql.sum(sales.units) jsql.multiply (products.price jsql.minus products.cost) jsql.as(profit)"
    ],
    "from": ["products"],
    "left_join": [
        "sales jsql.using product_id"
    ],
    "where": [
        "sales.date jsql.gt (CURRENT_DATE jsql.minus INTERVAL '4 weeks')"
    ],
    "group_by" : [
        "product_id"
        "products.name"
        "products.price"
        "products.cost"
    ],
    "having" :[
        "jsql.sum( products.price jsql.multiply sales.units ) jsql.gt 5000"
    ]
}
```

This is slightly less readable but likely easier to encode/decode.
worth trying to see if that's really a problem we need to fix or not...


With tokenisation taken care of and everything as a function, it's relatively simple
if we can cope with operators instead of just functions:


```json
{
    "select" :[
        "product_id",
        "products.name",
        "(sum(sales.units) * (products.price - products.cost) as profit"
    ],
    "from": ["products"],
    "left_join": [
        "sales using product_id"
    ],
    "where": [
        "sales.date > (CURRENT_DATE - INTERVAL '4 weeks')"
    ],
    "group_by" : [
        "products.product_id",
        "products.name",
        "products.price",
        "products.cost"
    ],
    "having" :[
        "sum( products.price * sales.units ) > 5000"
    ]
}
```

much more readable now, and relatively few operators added, still retaining function prefixes
if * / -> etc are not easily encoded/decoded by the users can try:
