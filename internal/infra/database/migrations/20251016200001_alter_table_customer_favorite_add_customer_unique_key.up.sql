ALTER TABLE customer_favorites
    ADD CONSTRAINT uk_customer_favorites_product_id UNIQUE (customer_id, product_id);