TODO:
========
- DB Query return types are available, add an option to make those available to the caller, this should enable runtime evaluation in client code and maintaining a cache of types in development for type safety, like Prisma.



Questions:
=============
- How to deal with malicious inserts, eg someone gets ahold of a form,
then proceeds to dump 10000 records per second into it...
recaptcha support is one method, how to support that?




In progress
============

- Add caching of database information and refresh the cache automatically every so often in a thread-safe way.
- Add validation in the query parsing against the API's cached schema information
- Adding specifications to the SQL query component fields, so we can identify insert, update, delete, select type of queries and check all fields are provided, non-empty, non-recursive (or limited recursion)
- Make 





Design ideas
===========

Syntax goal for "bootstrap" initial implementation
A mixture of:
    named arguments in expressions (eg select) where we need the flexibility
    positional arguments in more rigid parts, eg join and where and order_by

```json
{
    "select" :[
        { "table": "customers", "column": "name" }, 
        { "table": "customers", "column": "country" }
    ],
    "from": ["customers"],
    "left_join": [
        {
            "arg1": { "table": "orders", "column": "customer_id"}, 
            "op": "eq", 
            "arg2": { "table": "customers", "column": "id"}
        }
    ],
    "where": [
        {
            "left": { "table": "customers", "column": "name"}, 
            "op": "eq", 
            "right": { "table": "customers", "column": "id"}
        }
    ],
    "order_by" : [
        { "table": "customers", "column": "name", "order": "asc"}, 
        { "table": "customers", "column": "country", "order": "asc"}, 
    ]
}
```

A more complex SQL query to represent:

```sql
SELECT
	product_id,
	products.name,
	(SUM(sales.units) * (products.price - products.cost)) AS profit
FROM
	products
	LEFT JOIN sales on sales.product_id = products.product_id
WHERE
	sales.date > '2023-01-01'
GROUP BY
	products.product_id,
	products.name,
	products.price,
	products.cost
HAVING
	SUM(products.price * sales.units) > 5000;

```

More complex query is possible but it gets more nested
```json
{
    "select" :[
        { "table": "products", "column": "product_id" },
        { "table": "products", "column": "name" },
        { 
            "table": "sales", 
            "column": "units",
            "as": "profit",
            "fns": [
                { "fn":"sum" }, 
                { "fn": "mult", "args": [
                    { "table": "products", "column": "price", "fns": [
                        { 
                            "fn":"sub", "args": [ 
                                { "table": "products", "column":"price" } 
                            ] 
                        }
                    ]}
                ]}
            ] 
        }
    ],
    "from": ["products"],
    "left_join": [
        {
            "arg1": { "table": "sales", "column": "product_id" }, 
            "op": "eq", 
            "arg2": { "table": "products", "column": "product_id" }
        }
    ],
    "where": [
        {
            "left": { "table":"sales", "column":"date" }, 
            "op": "gt", 
            "right": { "values": ["2023-01-01"] }
        }
    ],
    "group_by" : [
        { "table": "products", "column": "product_id" },
        { "table": "products", "column": "name" },
        { "table": "products", "column": "price" },
        { "table": "products", "column": "cost" },
    ],
    "having" :[
        {
            "left": { 
                "table": "products", 
                "column": "price", 
                "fns": [
                    { 
                        "fn": "mult", "args": [ 
                            { "table": "sales", "column": "units" }
                        ]
                    },
                    { "fn": "sum" }
                ] 
            },
            "op": "gt",
            "right": { "values": ["5000"] }
        }
    ]
}
```

Once "bootstrap" is fully functional we can look to create a compiler to make this easier:
```json
{
    "select": [
        "product_id",
        "products.name",
        "(sum(sales.units) * (products.price - products.cost) as profit"
    ],
    "from": ["products"],
    "left_join": [
        "sales using product_id"
    ],
    "where": [
        "sales.date > '2023-01-01'"
    ],
    "group_by": [
        "products.product_id",
        "products.name",
        "products.price",
        "products.cost"
    ],
    "having": [
        "sum( products.price * sales.units ) > 5000"
    ]
}
```


