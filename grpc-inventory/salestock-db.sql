drop TABLE if EXISTS INVENTORY;
CREATE TABLE `inventory` (
        `sku` text PRIMARY KEY NOT NULL,
        `item_name` text NOT NULL,
        `stock` INTEGER NOT NULL,
        `created_on` datetime DEFAULT CURRENT_TIMESTAMP,
        `updated_on` datetime null
    );


drop TRIGGER if EXISTS timestamperupdateinventory;

CREATE TRIGGER timestamperupdateinventory
AFTER UPDATE ON INVENTORY
FOR EACH ROW
BEGIN
   UPDATE INVENTORY SET updated_on = datetime('now') WHERE sku = NEW.sku;
END;


drop TABLE if EXISTS PO_HEADERS;
CREATE TABLE `po_headers` (
        `po_header_id` INTEGER PRIMARY KEY AUTOINCREMENT,
        `po_date` datetime NOT NULL,
        `created_on` datetime DEFAULT CURRENT_TIMESTAMP,
        `updated_on` datetime null
    );

drop TRIGGER if EXISTS timestamperupdatepoheaders;

CREATE TRIGGER timestamperupdatepoheaders
AFTER UPDATE ON PO_HEADERS
FOR EACH ROW
BEGIN
   UPDATE PO_HEADERS SET updated_on = datetime('now') WHERE po_header_id = NEW.po_header_id;
END;


drop TABLE if EXISTS po_lines;
CREATE TABLE `po_lines` (
        `po_line_id` INTEGER PRIMARY KEY AUTOINCREMENT,
        `po_header_id` INTEGER NOT NULL,
        `sku` text NOT NULL,
        `pruchase_price` FLOAT NOT NULL,
        `qty_po` INTEGER NOT NULL,
        `status` text NOT NULL,
        `created_on` datetime DEFAULT CURRENT_TIMESTAMP,
        `updated_on` datetime null,
        FOREIGN KEY(po_header_id) REFERENCES PO_HEADERS(po_header_id),
        FOREIGN KEY(sku) REFERENCES inventory(sku)
    );

drop TRIGGER if EXISTS timestamperupdatepolines;

CREATE TRIGGER timestamperupdatepolines
AFTER UPDATE ON po_lines
FOR EACH ROW
BEGIN
   UPDATE po_lines SET updated_on = datetime('now') WHERE po_line_id = NEW.po_line_id;
END;

drop TABLE if EXISTS DELIVERY_ORDER;
CREATE TABLE `delivery_order` (
        `do_id` INTEGER PRIMARY KEY AUTOINCREMENT,
        `po_line_id` INTEGER NOT NULL,
        `sku` text NOT NULL,
        `kwitansi_num` text NOT NULL,
        `qty_ship` INTEGER NOT NULL,
        `status` text NOT NULL,
        `created_on` datetime DEFAULT CURRENT_TIMESTAMP,
        `updated_on` datetime null,
        FOREIGN KEY(po_line_id) REFERENCES po_lines(po_line_id),
        FOREIGN KEY(sku) REFERENCES inventory(sku)
    );

CREATE TRIGGER timestamperupdatedeliveryorder
AFTER UPDATE ON delivery_order
FOR EACH ROW
BEGIN
   UPDATE delivery_order SET updated_on = datetime('now') WHERE do_id = NEW.do_id;
END;

drop TABLE if EXISTS RECEIVING;
CREATE TABLE `receiving` (
        `receive_id` INTEGER PRIMARY KEY AUTOINCREMENT,
        `do_id` INTEGER NOT NULL,
        `receive_date` datetime NOT NULL,
        `qty_receive` INTEGER NOT NULL,
        `status` text NOT NULL,
        `created_on` datetime DEFAULT CURRENT_TIMESTAMP,
        `updated_on` datetime null,
        FOREIGN KEY(do_id) REFERENCES DELIVERY_ORDER(do_id)
    );

CREATE TRIGGER timestamperupdatereceiving
AFTER UPDATE ON receiving
FOR EACH ROW
BEGIN
   UPDATE receiving SET updated_on = datetime('now') WHERE receive_id = NEW.receive_id;
END;  


drop TABLE if EXISTS SALES_ORDER;
CREATE TABLE `sales_order` (
        `order_num` text PRIMARY KEY NOT NULL,
        `po_line_id` integer NOT NULL,
        `order_date` datetime NOT NULL,
        `created_on` datetime DEFAULT CURRENT_TIMESTAMP,
        `updated_on` datetime null,
        FOREIGN KEY(po_line_id) REFERENCES po_lines(po_line_id)
    );

CREATE TRIGGER timestamperupdatesalesorder
AFTER UPDATE ON sales_order
FOR EACH ROW
BEGIN
   UPDATE sales_order SET updated_on = datetime('now') WHERE order_num = NEW.order_num;
END;

drop TABLE if EXISTS SO_LINES;
CREATE TABLE `so_lines` (
        `so_line_id` INTEGER PRIMARY KEY AUTOINCREMENT,
        `order_num`  text NOT NULL,
        `sku` text NOT NULL,
        `selling_price` FLOAT NOT NULL,
        `qty_so` INTEGER NOT NULL,
        `status` text NOT NULL,
        `created_on` datetime DEFAULT CURRENT_TIMESTAMP,
        `updated_on` datetime null,
        FOREIGN KEY(order_num) REFERENCES SALES_ORDER(order_num),
        FOREIGN KEY(sku) REFERENCES inventory(sku)
    );


CREATE TRIGGER timestamperupdatesoline
AFTER UPDATE ON so_lines
FOR EACH ROW
BEGIN
   UPDATE so_lines SET updated_on = datetime('now') WHERE so_line_id = NEW.so_line_id;
END;

drop TABLE if EXISTS shipping;
CREATE TABLE `shipping` (
        `ship_id` INTEGER PRIMARY KEY AUTOINCREMENT,
        `so_line_id`  INTEGER NOT NULL,
        `sku` text NOT NULL,
        `shipping_date` datetime NOT NULL,
        `qty_ship` INTEGER NOT NULL,
        `status` text NOT NULL,
        `created_on` datetime DEFAULT CURRENT_TIMESTAMP,
        `updated_on` datetime null,
         FOREIGN KEY(so_line_id) REFERENCES SO_LINES(so_line_id)
    );

CREATE TRIGGER timestamperupdateshipping
AFTER UPDATE ON shipping
FOR EACH ROW
BEGIN
   UPDATE shipping SET updated_on = datetime('now') WHERE ship_id = NEW.ship_id;
END;