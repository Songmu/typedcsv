# CONCEPT

- Quoted fields are explicitly string
- Unquoted fields are guessed of types
    - null
    - number
        - integer
        - decimal(?)
        - float
    - boolean
        - true/false
    - date(?)
        - RFC3339

Unquoted fields that could not be type inferred will result in an error.
