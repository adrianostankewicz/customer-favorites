ALTER TABLE customer_favorites
    ADD CONSTRAINT fk_customer_favorites_customer
    FOREIGN KEY (customer_id)
    REFERENCES customer(id)
    ON DELETE CASCADE;