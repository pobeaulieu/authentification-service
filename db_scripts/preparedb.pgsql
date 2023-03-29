INSERT INTO clients (id, first_name, last_name, business_type, residential_type)
VALUES
    (1, 'John', 'Doe', 1, 0),
    (2, 'Jane', 'Doe', 1, 1),
    (3, 'Bob', 'Smith', 0, 1),
    (4, 'Alice', 'Johnson', 1, 0),
    (5, 'Mike', 'Williams', 0, 0),
    (6, 'Sarah', 'Davis', 0, 1),
    (7, 'Tom', 'Brown', 1, 0),
    (8, 'Mary', 'Miller', 0, 0),
    (9, 'David', 'Wilson', 1, 1),
    (10, 'Karen', 'Taylor', 0, 1);

INSERT INTO password_policies (min_length, require_upper, require_lower, require_number, require_symbol, history_size) 
VALUES (4, false, true, true, false, 2);

INSERT INTO login_policies (max_login_attempt_count, login_time_interval) 
VALUES (5,5);

