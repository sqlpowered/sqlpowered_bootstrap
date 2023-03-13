# Project Goals


> <br>
> <p style="color:red"><b>Pre-Alpha software, still a work in progress</b></p>
> 
> <br>
<br>

> <br>
> Make Postgres into a "database as a service", with a Dynamic, scalable, safe API to the Postgres database.
> 
> <br>
<br>

## Simple
- Queries can respond to user input in real time, ideal for web **frontend** or **mobile** development, can change the query structure in real time. No need to make many different APIs.
- Write queries in JSON, which looks like SQL. If you know SQL you already know SQLPowered.
- No need for an ORM, edit your queries as JSON.

## Supported
- Simple queries are easy, but still supports complex queries with common table expressions, case conditional logic, advanced filtering, grouping, etc.
- API is supported everywhere with an internet connection, no need for a custom client, including web frontends and backends, the edge, embedded, mobile and in other untrusted environments.

## Secure
- Roles based access management.
- Roles control access to each table and column for each SQL operation, including what can be filtered on.
- Restrict a customer to only see their own data automatically. 
- SQL injection is impossible.

## Managed
- API rate limiting.
- Query cost budget, budget for the number of rows, columns, number of joins and query recursion depth.


# Example queries

I want customer `name` and `country`
```json
{
    "select": [
        { "table": "customers", "column": "name" }, 
        { "table": "customers", "column": "country" }
    ],
    "from": ["customers"]
}
```



I want to look at a specific customer's order items
```json
{
    "select": [
        { "table": "customers", "column": "name" }, 
        { "table": "customers", "column": "country" }
    ],
    "from": ["customers"],
    "left_join": [
        {
            "arg1": { "table": "orders", "column": "customer_id" }, 
            "op": "eq", 
            "arg2": { "table": "customers", "column": "id" }
        }
    ],
    "where": [
        {
            "left": { "table": "customers", "column": "name" }, 
            "op": "eq", 
            "right": { "table": "customers", "column": "id" }
        }
    ],
    "order_by" : [
        { "table": "customers", "column": "name", "order": "asc" }, 
        { "table": "customers", "column": "country", "order": "asc" }, 
    ]
}
```



TODO:
========
- Add query cache, with tiered caching, akin to CPU cache structure, but inverted,
where repeatedly accessed queries are promoted to higher priority caches
- valid sql check and generation -- this can be cached separetly from...
- valid sql permissions for particular user -- cache user's permissions


Questions:
=============
- How to deal with malicious inserts, eg someone gets ahold of a form,
then proceeds to dump 10000 records per second into it... recaptcha support is one method
- DB Query return types are available, add an option to make those available to the caller, this should enable runtime evaluation in client code and maintaining a cache of types in development for type safety, like Prisma.



Progress towards our goals:
============

| Task                                              | Progress |
|---------------------------------------------------|----------|
| Define query syntax and matching data types. | 90% |
| Document API goals, scope, audience and capabilities. | 0% |
| Example docker compose to start DB | 100% |
| Example data in SQL | 100% |
| Example API application running in docker compose | 0% |
| Processing query response data into: JSON key value pairs, JSON arrays, CSV  | 75% |
| Basic config management in JSON | 100% |
| Basic feature set, valid values (functions, types operators etc) | 100% |
| SQL component query spec, to identify each type of insert, update, delete, select define a minimum set of fields is supplied, validate allowed recurision levels. | 10% |
| Obtaining db info | 40% |
| Naive SQL builder: select | 0% |
| Naive SQL builder: select functions | 0% |
| Naive SQL builder: from | 0% |
| Naive SQL builder: where/having | 0% |
| Naive SQL builder: join | 0% |
| Naive SQL builder: group_by | 0% |
| Naive SQL builder: order_by | 0% |
| Naive SQL builder: cte | 0% |
| Query validation: select | 5% |
| Query validation: select functions | 0% |
| Query validation: from | 0% |
| Query validation: where/having | 0% |
| Query validation: join | 0% |
| Query validation: group_by | 0% |
| Query validation: order_by | 0% |
| Query validation: cte | 0% |
| Query validation: against cached db info | 0% |
| Thread-safe Caching: db info | 0% |
| Thread-safe Caching: queries | 0% |

