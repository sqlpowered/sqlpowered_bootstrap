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





