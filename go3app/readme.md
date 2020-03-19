Service 1. Additional, send JSON.

1. Create data string.
2. Print data on the screen.
2. Encode data string into JSON.
3. Send JSON to service 2 on port http :8082
_____________

Service 2. MAIN SERVICE (NUMBER 3 FOR DEMO LAB). ENCODING JSON->YAML.

1. Listen incoming traffic on port http :8082
2. Decode structure from JSON into map.
3. Print mapped data on the screen.
4. Encode data map into YAML.
5. Send YAML to service 3 on port http :8083
_____________

Service 3. Additional, get YAML.

1. Listen incoming traffic on port http :8083
2. Decode structure from YAML into map.
3. Print mapped data on the screen.
________

* Database structure with example:

emp_id: '1' (Employees_ID)

first_name: Bill

second_name: Gates

types: developer   (developer, designer, manager)

experience: '10'

default_salary: '2000'

________
