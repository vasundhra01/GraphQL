# GraphQL 
## For Elmeasure

---

# 1. GraphQL Introduction

GraphQL is a query language and runtime for APIs.
GraphQL is not a database.
It is mainly an API layer between frontend apps and backend services/databases.  
It allows the frontend to request exactly the data it needs from the backend.

typeDef: every piece of data, query defined in graphQL sits inside typedef  
resolvers: all the functions that resolve the types and make calls to API, calls databses  

---
# 2. Why GraphQL Was Created

Traditional REST APIs caused problems:

- Too many API endpoints
- Too many requests increase frontend load and more bandwidth used
- Overfetching unnecessary data. Underfetching imp data
- Frontend/backend coordination complexity
- increased latency

Example:

```txt
GET /plants/1
GET /plants/1/meters
GET /meters/11/alerts
GET /meters/11/readings
```
### REST 
```txt
GET /meter/11

Returns:

{
  "id": 11,
  "name": "Main Meter",
  "serialNo": "ABC123",
  "manufacturer": "XYZ",
  "location": "Floor 2",
  "latestReading": 420
}
```
Extra data is wasted if only Id and name rquired  
Frontend may need many requests for one screen.

GraphQL solves this by allowing one structured query.

---

# 3. Core Idea of GraphQL

Frontend specifies:
- what data it wants
- which fields it needs
- how deeply nested relationships should go

Backend returns exactly that structure.  
GraphQL allows to traverse connected data structures dynamically
```txt
Organization
 └── Plants
      └── Buildings
           └── Meters
                └── Readings
                     └── Alerts
```
Everything is related.

Traditional REST treats these as:

+ separate endpoints
+ separate requests
+ separate resources

GraphQL treats them as:
connected traversable objects  
Frontend walks through connected relationships naturally. Unlike rest where frontend combines data manually

# 4. REST vs GraphQL

| REST | GraphQL |
|---|---|
| Many endpoints | Single endpoint |
| Backend decides response shape | Frontend decides |
| Overfetching common | Exact fields fetched |
| Easier initially | More flexible |
| Simple caching | Harder caching |
| Good for simple APIs | Good for connected systems |

---

# 5. Example — REST vs GraphQL

## REST

```txt
GET /plant/1
GET /plant/1/meters
GET /meter/11/latest
GET /meter/11/alerts
```

---

## GraphQL

```graphql
query {
  plant(id: 1) {
    name

    meters {
      name

      latestReading {
        value
      }

      alerts {
        severity
      }
    }
  }
}
```
---

# 7. Main Components of GraphQL

---

## 7.1 Schema

Defines:
- available data
- relationships
- data types/ field type
- opeartions allowed
- api strcuture

Example:

```graphql
type Plant {
  id: ID!
  name: String!
  meters: [Meter]
}
type Meter {
  id: ID!
  name: String!
  latestReading: Float
}
```
``` Meaning:

Plant has:
- id
- name
- list of meters -- here meters are connected to plants so frontend can naturally traverse

Meter has:
- id
- name
```
In REST:  
URL defines operation-> backend decides  
In GraphQL:  
query defines operation-> frontend decides  

## 7.2 Query
Used to fetch/read data.

Example:

```graphql
query {
  plant(id: 1) {
    name
  }
}
```
Frontend can traverse all relationships in ONE query.
Traversal is nested. No need to call multiple API or manually combine data. async req can be managed.  
Prob: Nested traversal can become computationally expensive and can cause performance engineering problem
### BATCHING
GraphQL systems use batching tools like DataLoader
Instead of:
```
1 query → get all plants
Then:
100 separate queries → get meters for each plant
100 queries for meters assuming there are a 100 meters
```
Do:
``` txt
1 query:
SELECT * FROM meters
WHERE plant_id IN (...)
```
Huge optimization.

GraphQL queries LOOK simple.
But backend execution may become extremely complex.
This is why GraphQL is harder than REST at scale.

## 7.3 Mutation

Used to modify/change data.

Equivalent to:
```txt
POST
PUT
DELETE
```

Example:

```graphql
mutation {
  createMeter(name: "Boiler Meter") {
    id
    name
  }
}
```

---
## 7.4 SUBSCRIPTIONS — REAL-TIME GRAPHQL
enables:
#### live updates
Example:
+ live meter readings
+ live alert notifications
+ dashboard auto-refresh
keeps sending updates whenever meter changes-- connection stays open, uses
#### WebSockets
---
## 7.5 Resolver

Resolver is backend logic that fetches data.
Resolver may:
- query database
- call another service
- perform calculations

Resolvers power GraphQL internally.

Suppose database:
``` txt

Plant Table
Meter Table
```
Resolver logic:
```txt

Find meters where plant_id = current plant
Resolvers connect schema to real data.
```
---

# 8. How GraphQL Works Internally

```txt
Frontend
    ↓
GraphQL Query
    ↓
GraphQL Server (reads wuery structure)
    ↓
Schema Validation (checks if it exists)
    ↓
Resolvers Execute (resolver fetches actual data)
    ↓
Database / Services
    ↓
JSON Response Returned (query shape determines response shape)
```

---

# 9. GraphQL Response Structure

Example:

```json
{
  "data": {
    "plant": {
      "name": "Bangalore Plant"
    }
  }
}
```

GraphQL responses are usually tree-shaped JSON.

---

# 10. Context in GraphQL

Each request carries contextual info such as permissions, current user, token etc.  
Resolver uses context

Example:

```graphql
query {
  plant(id: 1) {
    name

    meters {
      name

      latestReading {
        value
      }
    }
  }
}
```

The reading automatically comes with context:
- which plant
- which meter
- hierarchy relationships

Frontend receives connected information instead of isolated records.
Resolver checks eg Which organization does user belong to? then filters data  
Context is imp for authentication (who are you) and authorization (what are you allowed to access)
Gives role based access 


# 11. Overfetching and Underfetching

---

## Overfetching

Fetching unnecessary data.

Example:

```json
{
  "id": 11,
  "name": "Meter",
  "serialNo": "ABC",
  "manufacturer": "XYZ",
  "latestReading": 420
}
```

Frontend only needed:
```txt
name + latestReading
```

---

## Underfetching

Not getting enough data in one request.

Requires multiple API calls.

GraphQL solves both problems.

---

# 12. Why Frontend Developers Like GraphQL

Frontend controls:
- fields
- nesting
- response structure

Backend no longer needs:
- custom dashboard endpoints
- many specialized APIs

Frontend becomes more flexible.

---

# 13. GraphQL and Databases

```txt
Frontend
    ↓
GraphQL API Layer
    ↓
Database / Services
```
GraphQL is Frontend query orchestration layer and not storage engine. Resolvers internally use SQL, MomgoDB, REST calls eyc. 

# 14. GraphQL in Elmeasure

Elmeasure data is naturally hierarchical:

```txt
Organization
 └── Plant
      └── Building
           └── Meter
                └── Reading
                     └── Alert
```

GraphQL is excellent for:
- dashboards
- enterprise hierarchy browsing
- reports
- customer portals
- analytics summaries


# Good Use Cases in Elmeasure

---

## Dashboard APIs

Example:

```graphql
query {
  plant(id: 1) {
    energyToday

    meters {
      latestReading
    }
  }
}
```

---

## Multi-tenant Enterprise Systems

Useful when:
- many organizations
- many plants
- deeply connected data

---

## Custom Dashboards

Different clients need different metrics.

GraphQL lets frontend request only relevant fields.

---

## Mobile Apps

Reduces unnecessary payload size.

---

## API Gateway for Microservices

GraphQL can unify:
- EMS service
- Alert service
- Analytics service
- Billing service

into one API.

---

# 17. Where GraphQL is BAD

---

## High-Speed IoT Ingestion

Bad fit for:
- millions of readings
- telemetry streams
- real-time sensor ingestion



---

## Heavy Historical Analytics

Example:

```txt
5 years of second-by-second readings
```

Can become inefficient.

Better:
- analytics engines
- batch pipelines
- pre-aggregations

---

## Tiny Simple Applications

REST is simpler for:
- login
- health checks
- basic CRUD

---

# 18. Recommended Elmeasure Architecture

```txt
IoT Devices
    ↓
Protocol/ kafka - real time data pipeline
    ↓
Ingestion Services
    ↓
Database
    ↓
Business Logic Services
    ↓
GraphQL API Layer
    ↓
Frontend Dashboard
```

---

# 19. Important GraphQL Concepts

---

## Strong Typing

GraphQL schemas are typed.

Example:

```graphql
type Meter {
  latestReading: Float
}
```

Benefits:
- safer APIs
- better developer tooling
- autocomplete
- validation

---

## Nested Queries

GraphQL naturally supports hierarchy.

Example:

```graphql
query {
  company(id: 1) {
    plants {
      meters {
        alerts {
          severity
        }
      }
    }
  }
}
```

---

## Single Endpoint

Usually:

```txt
POST /graphql
```

Unlike REST:
- many endpoints not needed
- the operation is inside the query not the url

---

## 20. GraphQL Does NOT Automatically Do AI

GraphQL only follows relationships you define.

It does NOT infer meaning.

It cannot automatically conclude:

```txt
High energy usage means motor inefficiency
```

That requires:
- analytics
- AI
- business logic

---

## 21. GraphQL Security Concerns

GraphQL can be dangerous if unmanaged.

---

## Deep Query Attacks

Client may request massive nested data.

Example:

```txt
company → plants → meters → readings → alerts
```

Need:
- depth limits
- query cost limits

---

## Expensive Queries

Large nested queries can overload backend.

Need:
- optimization
- pagination
- caching

---

## Pagination
Prevents huge responses.

```graphql
query {
  meters(limit: 10, offset: 0) {
    name
  }
}
```
GraphQL allows flexible querying.
Without limits:
clients can accidentally destroy performance.

---
## Query Complexity problem
Greater flexibility could cause massive joins, huge memory usage and latency.
Solution: Production systems enforces-
+ query depth limits (set max nesting limit to 5- prevents malicious deep query)
+ field limits (pagination)
+ execution timeouts
+ complexity scoring
---
## Subscriptions (Real-Time GraphQL)

GraphQL supports real-time updates.

Example:
- live meter readings
- alert notifications

Uses:
- WebSockets

Example:

```graphql
subscription {
  meterUpdated(id: 11) {
    latestReading
  }
}
```

Useful for dashboards.

---


## 24. Final Recommendation (for Elmeasure)

Use GraphQL for:
- dashboards
- enterprise hierarchy browsing
- reports
- mobile apps
- customer portals
- aggregated APIs
- request-respons efocused
good for schema parsing, query parsing, resolver execution, traversal logic

Do NOT use GraphQL for:
- telemetry ingestion (cont., massive, high freq)
- streaming pipelines - sensors cont. generate data, system must process real time and tolerate bursts
- raw industrial event processing
GraphQL was designed for frontend querying flexibility not fast ingestion. for that use MQTT, Kafka, raw REST ingestion API that re optimized for throughput, streaming and low latency
Best architecture:

```txt
Kafka/MQTT + Time-Series Database + GraphQL API Layer
```
## GRAPHQL AS ORCHESTRATION LAYER

Frontend sends ONE query:
```
query {
  plant(id: 1) {
    alerts
    analytics
    billing
  }
}
```
GraphQL internally orchestrates:
+ alert service
+ analytics service
+ billing service
then combines response.

Frontend sees one unified graph
#### Without orchestration:
Frontend manually calls:
service A
service B
service C
#### GraphQL centralizes orchestration.
