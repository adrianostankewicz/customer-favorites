CREATE TABLE IF NOT EXISTS customer_favorites (
  id UUID NOT NULL,
  product_title VARCHAR(100) NOT NULL,
  product_id BIGINT NOT NULL,
  product_image VARCHAR(255) NOT NULL,
  product_price BIGINT NOT NULL,
  product_review_rate INTEGER NULL,
  product_review_count INTEGER NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP NULL
);