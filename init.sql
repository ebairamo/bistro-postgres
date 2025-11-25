-- ENUM Types
CREATE TYPE order_status AS ENUM ('open', 'close');
CREATE TYPE unit AS ENUM ('ml', 'g', 'shots');

-- INVENTORY TABLE
CREATE TABLE inventory (
  id SERIAL PRIMARY KEY,
  ingredient_id VARCHAR(100) UNIQUE NOT NULL,
  name VARCHAR(100) NOT NULL,
  quantity FLOAT NOT NULL,
  unit unit NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- MENU_ITEMS TABLE
CREATE TABLE menu_items (
  id SERIAL PRIMARY KEY,
  product_id VARCHAR(100) UNIQUE NOT NULL,
  name VARCHAR(100) NOT NULL,
  description VARCHAR(255) NOT NULL,
  price FLOAT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- MENU_ITEM_INGREDIENTS TABLE (связь между меню и ингредиентами)
CREATE TABLE menu_item_ingredients (
  id SERIAL PRIMARY KEY,
  menu_item_id INT NOT NULL,
  ingredient_id INT NOT NULL,
  quantity FLOAT NOT NULL,
  FOREIGN KEY (menu_item_id) REFERENCES menu_items(id),
  FOREIGN KEY (ingredient_id) REFERENCES inventory(id)
);

-- ORDERS TABLE
CREATE TABLE orders (
  id SERIAL PRIMARY KEY,
  order_id VARCHAR(100) UNIQUE NOT NULL,
  customer_name VARCHAR(100) NOT NULL,
  status order_status NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- ORDER_ITEMS TABLE (связь между заказами и меню)
CREATE TABLE order_items (
  id SERIAL PRIMARY KEY,
  order_id INT NOT NULL,
  menu_item_id INT NOT NULL,
  quantity INT NOT NULL,
  FOREIGN KEY (order_id) REFERENCES orders(id),
  FOREIGN KEY (menu_item_id) REFERENCES menu_items(id)
);