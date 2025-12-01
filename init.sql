-- ===== ENUM TYPES =====
CREATE TYPE order_status AS ENUM ('open', 'close');
CREATE TYPE unit AS ENUM ('ml', 'g', 'shots');
CREATE TYPE payment_method AS ENUM ('cash', 'card', 'online');

-- ===== TABLES =====

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

-- MENU_ITEM_INGREDIENTS TABLE (Junction Table)
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
  total_amount FLOAT DEFAULT 0,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- ORDER_ITEMS TABLE (Junction Table)
CREATE TABLE order_items (
  id SERIAL PRIMARY KEY,
  order_id INT NOT NULL,
  menu_item_id INT NOT NULL,
  quantity INT NOT NULL,
  price_at_order FLOAT NOT NULL,
  FOREIGN KEY (order_id) REFERENCES orders(id),
  FOREIGN KEY (menu_item_id) REFERENCES menu_items(id)
);

-- ORDER_STATUS_HISTORY TABLE
CREATE TABLE order_status_history (
  id SERIAL PRIMARY KEY,
  order_id INT NOT NULL,
  old_status order_status,
  new_status order_status NOT NULL,
  changed_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (order_id) REFERENCES orders(id)
);

-- PRICE_HISTORY TABLE
CREATE TABLE price_history (
  id SERIAL PRIMARY KEY,
  menu_item_id INT NOT NULL,
  old_price FLOAT,
  new_price FLOAT NOT NULL,
  changed_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (menu_item_id) REFERENCES menu_items(id)
);

-- INVENTORY_TRANSACTIONS TABLE
CREATE TABLE inventory_transactions (
  id SERIAL PRIMARY KEY,
  ingredient_id INT NOT NULL,
  transaction_type VARCHAR(50) NOT NULL, -- 'add', 'remove', 'adjust'
  quantity_change FLOAT NOT NULL,
  old_quantity FLOAT,
  new_quantity FLOAT,
  reason VARCHAR(255),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (ingredient_id) REFERENCES inventory(id)
);

-- ===== INDEXES =====

-- Inventory indexes
CREATE INDEX idx_inventory_ingredient_id ON inventory(ingredient_id);
CREATE INDEX idx_inventory_created_at ON inventory(created_at);

-- Menu items indexes
CREATE INDEX idx_menu_items_product_id ON menu_items(product_id);
CREATE INDEX idx_menu_items_price ON menu_items(price);
CREATE INDEX idx_menu_items_created_at ON menu_items(created_at);

-- Menu item ingredients indexes
CREATE INDEX idx_menu_item_ingredients_menu_item_id ON menu_item_ingredients(menu_item_id);
CREATE INDEX idx_menu_item_ingredients_ingredient_id ON menu_item_ingredients(ingredient_id);

-- Orders indexes
CREATE INDEX idx_orders_order_id ON orders(order_id);
CREATE INDEX idx_orders_customer_name ON orders(customer_name);
CREATE INDEX idx_orders_status ON orders(status);
CREATE INDEX idx_orders_created_at ON orders(created_at);

-- Order items indexes
CREATE INDEX idx_order_items_order_id ON order_items(order_id);
CREATE INDEX idx_order_items_menu_item_id ON order_items(menu_item_id);

-- Order status history indexes
CREATE INDEX idx_order_status_history_order_id ON order_status_history(order_id);
CREATE INDEX idx_order_status_history_changed_at ON order_status_history(changed_at);

-- Price history indexes
CREATE INDEX idx_price_history_menu_item_id ON price_history(menu_item_id);
CREATE INDEX idx_price_history_changed_at ON price_history(changed_at);

-- Inventory transactions indexes
CREATE INDEX idx_inventory_transactions_ingredient_id ON inventory_transactions(ingredient_id);
CREATE INDEX idx_inventory_transactions_created_at ON inventory_transactions(created_at);

-- ===== TEST DATA =====

-- Inventory Items (20+ items)
INSERT INTO inventory (ingredient_id, name, quantity, unit) VALUES
('espresso_shot', 'Espresso Shot', 500, 'shots'),
('sugar', 'Sugar', 5000, 'g'),
('flour', 'Flour', 10000, 'g'),
('milk', 'Milk', 20000, 'ml'),
('blueberries', 'Blueberries', 2000, 'g'),
('blueberriesVip', 'Blueberries VIP', 2500, 'g'),
('water', 'Water', 50000, 'ml'),
('chocolate', 'Chocolate', 1500, 'g'),
('vanilla', 'Vanilla Extract', 500, 'ml'),
('cinnamon', 'Cinnamon', 300, 'g'),
('honey', 'Honey', 800, 'ml'),
('butter', 'Butter', 2000, 'g'),
('eggs', 'Eggs', 100, 'shots'),
('baking_powder', 'Baking Powder', 500, 'g'),
('salt', 'Salt', 1000, 'g'),
('cream', 'Cream', 5000, 'ml'),
('coffee_beans', 'Coffee Beans', 3000, 'g'),
('tea', 'Tea Leaves', 800, 'g'),
('caramel', 'Caramel Sauce', 1000, 'ml'),
('hazelnut', 'Hazelnut', 1200, 'g'),
('almond', 'Almond Milk', 3000, 'ml');

-- Menu Items (10+ items)
INSERT INTO menu_items (product_id, name, description, price) VALUES
('latte', 'Caffe Latte', 'Espresso with steamed milk', 3.5),
('muffin', 'Blueberry Muffin', 'Freshly baked muffin with blueberries', 2.0),
('espresso', 'Espresso', 'Strong and bold coffee', 2.5),
('cappuccino', 'Cappuccino', 'Espresso with foam', 4.0),
('americano', 'Americano', 'Espresso with hot water', 2.75),
('mocha', 'Mocha', 'Espresso with chocolate and milk', 4.5),
('macchiato', 'Macchiato', 'Espresso with a splash of milk', 3.25),
('croissant', 'Croissant', 'Buttery pastry', 2.5),
('cheesecake', 'Cheesecake', 'Creamy cheesecake', 4.75),
('chocolate_cake', 'Chocolate Cake', 'Rich chocolate cake', 3.75),
('vanilla_latte', 'Vanilla Latte', 'Latte with vanilla', 3.75),
('honey_coffee', 'Honey Coffee', 'Coffee with honey', 3.5);

-- Menu Item Ingredients (Recipes)
INSERT INTO menu_item_ingredients (menu_item_id, ingredient_id, quantity) VALUES
-- Latte: espresso, milk
(1, 1, 1),
(1, 4, 200),
-- Muffin: flour, blueberries, sugar, eggs, butter
(2, 3, 100),
(2, 5, 20),
(2, 2, 30),
(2, 13, 2),
(2, 12, 50),
-- Espresso: espresso
(3, 1, 1),
-- Cappuccino: espresso, milk, cream
(4, 1, 1),
(4, 4, 150),
(4, 16, 50),
-- Americano: espresso, water
(5, 1, 1),
(5, 7, 200),
-- Mocha: espresso, milk, chocolate
(6, 1, 2),
(6, 4, 150),
(6, 8, 30),
-- Macchiato: espresso, milk
(7, 1, 1),
(7, 4, 50),
-- Croissant: flour, butter, salt
(8, 3, 150),
(8, 12, 100),
(8, 15, 5),
-- Cheesecake: flour, eggs, butter, sugar
(9, 3, 200),
(9, 13, 6),
(9, 12, 150),
(9, 2, 100),
-- Chocolate Cake: flour, chocolate, eggs, butter, sugar
(10, 3, 250),
(10, 8, 100),
(10, 13, 8),
(10, 12, 200),
(10, 2, 150),
-- Vanilla Latte: espresso, milk, vanilla
(11, 1, 1),
(11, 4, 200),
(11, 9, 5),
-- Honey Coffee: espresso, water, honey
(12, 1, 2),
(12, 7, 200),
(12, 11, 20);

-- Orders (30+ orders with varied statuses)
INSERT INTO orders (order_id, customer_name, status, total_amount) VALUES
('order001', 'Alice Smith', 'close', 7.5),
('order002', 'Bob Johnson', 'close', 2.5),
('order003', 'Charlie Brown', 'open', 5.5),
('order004', 'Diana Prince', 'close', 8.75),
('order005', 'Evan Davis', 'close', 3.5),
('order006', 'Fiona Green', 'open', 4.0),
('order007', 'George Harris', 'close', 6.25),
('order008', 'Helen White', 'close', 2.0),
('order009', 'Ian Black', 'open', 7.0),
('order010', 'Julia King', 'close', 9.5),
('order011', 'Kevin Lee', 'open', 4.5),
('order012', 'Laura Martin', 'close', 3.75),
('order013', 'Michael Clark', 'close', 5.5),
('order014', 'Nancy Wilson', 'open', 6.0),
('order015', 'Oscar Moore', 'close', 2.75),
('order016', 'Patricia Taylor', 'close', 8.0),
('order017', 'Quinn Adams', 'open', 4.25),
('order018', 'Rachel Scott', 'close', 7.5),
('order019', 'Samuel Green', 'open', 3.5),
('order020', 'Tina Baker', 'close', 5.0),
('order021', 'Ulysses Hill', 'close', 9.25),
('order022', 'Violet Allen', 'open', 4.0),
('order023', 'William Young', 'close', 6.5),
('order024', 'Xena Hall', 'open', 2.25),
('order025', 'Yuri Nelson', 'close', 7.75),
('order026', 'Zoe Carter', 'close', 3.5),
('order027', 'Aaron Mitchell', 'open', 5.25),
('order028', 'Bella Perez', 'close', 8.5),
('order029', 'Caleb Roberts', 'open', 4.75),
('order030', 'Daisy Phillips', 'close', 6.0),
('order031', 'Ethan Campbell', 'open', 3.25),
('order032', 'Faye Parker', 'close', 9.0);

-- Order Items
INSERT INTO order_items (order_id, menu_item_id, quantity, price_at_order) VALUES
-- order001: 2x latte, 1x muffin
(1, 1, 2, 3.5),
(1, 2, 1, 2.0),
-- order002: 1x espresso
(2, 3, 1, 2.5),
-- order003: 1x cappuccino, 1x croissant
(3, 4, 1, 4.0),
(3, 8, 1, 2.5),
-- order004: 1x mocha, 1x cheesecake
(4, 6, 1, 4.5),
(4, 9, 1, 4.75),
-- order005: 1x latte, 1x espresso
(5, 1, 1, 3.5),
(5, 3, 1, 2.5),
-- order006: 2x americano
(6, 5, 2, 2.75),
-- order007: 1x macchiato, 1x chocolate_cake
(7, 7, 1, 3.25),
(7, 10, 1, 3.75),
-- order008: 1x espresso
(8, 3, 1, 2.0),
-- order009: 3x latte
(9, 1, 3, 3.5),
-- order010: 2x mocha, 1x cheesecake
(10, 6, 2, 4.5),
(10, 9, 1, 4.75),
-- order011: 1x cappuccino, 1x muffin
(11, 4, 1, 4.0),
(11, 2, 1, 2.5),
-- order012: 1x vanilla_latte, 1x croissant
(12, 11, 1, 3.75),
(12, 8, 1, 2.5),
-- order013: 2x espresso, 1x honey_coffee
(13, 3, 2, 2.5),
(13, 12, 1, 3.5),
-- order014: 1x cappuccino
(14, 4, 1, 4.0),
-- order015: 1x americano, 1x espresso
(15, 5, 1, 2.75),
(15, 3, 1, 2.5),
-- order016: 1x chocolate_cake, 1x cheesecake
(16, 10, 1, 3.75),
(16, 9, 1, 4.75),
-- order017: 2x latte
(17, 1, 2, 3.5),
-- order018: 1x mocha, 1x croissant
(18, 6, 1, 4.5),
(18, 8, 1, 2.75),
-- order019: 1x macchiato
(19, 7, 1, 3.5),
-- order020: 1x cappuccino, 1x muffin
(20, 4, 1, 4.0),
(20, 2, 1, 2.0),
-- order021: 2x espresso, 1x chocolate_cake
(21, 3, 2, 2.5),
(21, 10, 1, 4.75),
-- order022: 1x vanilla_latte, 1x croissant
(22, 11, 1, 3.75),
(22, 8, 1, 2.5),
-- order023: 1x americano, 2x espresso
(23, 5, 1, 2.75),
(23, 3, 2, 2.5),
-- order024: 1x latte
(24, 1, 1, 3.5),
-- order025: 1x mocha, 1x cheesecake
(25, 6, 1, 4.5),
(25, 9, 1, 4.75),
-- order026: 1x cappuccino, 1x muffin
(26, 4, 1, 4.0),
(26, 2, 1, 2.5),
-- order027: 2x latte, 1x espresso
(27, 1, 2, 3.5),
(27, 3, 1, 2.5),
-- order028: 1x chocolate_cake, 1x cheesecake
(28, 10, 1, 3.75),
(28, 9, 1, 4.75),
-- order029: 1x cappuccino, 1x croissant
(29, 4, 1, 4.0),
(29, 8, 1, 2.5),
-- order030: 1x macchiato, 1x muffin
(30, 7, 1, 3.25),
(30, 2, 1, 2.75),
-- order031: 1x latte
(31, 1, 1, 3.5),
-- order032: 2x mocha
(32, 6, 2, 4.5);

-- Order Status History
INSERT INTO order_status_history (order_id, old_status, new_status) VALUES
(1, NULL, 'open'),
(1, 'open', 'close'),
(2, NULL, 'open'),
(2, 'open', 'close'),
(3, NULL, 'open'),
(4, NULL, 'open'),
(4, 'open', 'close'),
(5, NULL, 'open'),
(5, 'open', 'close'),
(6, NULL, 'open'),
(7, NULL, 'open'),
(7, 'open', 'close'),
(8, NULL, 'open'),
(8, 'open', 'close'),
(9, NULL, 'open'),
(10, NULL, 'open'),
(10, 'open', 'close');

-- Price History (tracking price changes)
INSERT INTO price_history (menu_item_id, old_price, new_price) VALUES
(1, 3.25, 3.5),
(2, 1.75, 2.0),
(4, 3.75, 4.0),
(6, 4.25, 4.5),
(10, 3.5, 3.75);

-- Inventory Transactions
INSERT INTO inventory_transactions (ingredient_id, transaction_type, quantity_change, old_quantity, new_quantity, reason) VALUES
(1, 'remove', -50, 500, 450, 'Used in orders'),
(4, 'remove', -100, 20000, 19900, 'Used in orders'),
(5, 'remove', -30, 2000, 1970, 'Used in muffins'),
(3, 'remove', -200, 10000, 9800, 'Used in baking'),
(1, 'add', 100, 450, 550, 'Restocked'),
(4, 'add', 500, 19900, 20400, 'Restocked'),
(8, 'remove', -50, 1500, 1450, 'Used in mocha'),
(12, 'remove', -100, 2000, 1900, 'Used in baking'),
(13, 'remove', -20, 100, 80, 'Used in recipes');