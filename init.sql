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

-- MENU_ITEM_INGREDIENTS TABLE
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

-- ORDER_ITEMS TABLE
CREATE TABLE order_items (
  id SERIAL PRIMARY KEY,
  order_id INT NOT NULL,
  menu_item_id INT NOT NULL,
  quantity INT NOT NULL,
  FOREIGN KEY (order_id) REFERENCES orders(id),
  FOREIGN KEY (menu_item_id) REFERENCES menu_items(id)
);

-- ===== TEST DATA =====

-- Inventory Items
INSERT INTO inventory (ingredient_id, name, quantity, unit) VALUES
('espresso_shot', 'Espresso Shot', 474, 'shots'),
('sugar', 'Sugar', 4790, 'g'),
('flour', 'Flour', 9300, 'g'),
('milk', 'Milk', 16000, 'ml'),
('blueberries', 'Blueberries', 1860, 'g'),
('blueberriesVip', 'Blueberries VIP', 2000, 'g');

-- Menu Items
INSERT INTO menu_items (product_id, name, description, price) VALUES
('latte', 'Caffe Latte', 'Espresso with steamed milk', 3.5),
('muffin', 'Blueberry Muffin', 'Freshly baked muffin with blueberries', 2.0),
('espresso', 'Espresso', 'Strong and bold coffee', 2.5);

-- Menu Item Ingredients (рецепты)
INSERT INTO menu_item_ingredients (menu_item_id, ingredient_id, quantity) VALUES
(1, 1, 1),      -- latte: 1 espresso_shot
(1, 4, 200),    -- latte: 200 ml milk
(2, 3, 100),    -- muffin: 100g flour
(2, 5, 20),     -- muffin: 20g blueberries
(2, 2, 30),     -- muffin: 30g sugar
(3, 1, 1);      -- espresso: 1 espresso_shot

-- Orders
INSERT INTO orders (order_id, customer_name, status) VALUES
('order123', 'Alice Smith', 'open'),
('order456', 'Bob Johnson', 'close'),
('order789', 'Charlie Brown', 'open');

-- Order Items
INSERT INTO order_items (order_id, menu_item_id, quantity) VALUES
(1, 1, 2),      -- order123: 2x latte
(1, 2, 1),      -- order123: 1x muffin
(2, 3, 1),      -- order456: 1x espresso
(3, 1, 1);      -- order789: 1x latte