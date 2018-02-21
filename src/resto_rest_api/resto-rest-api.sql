CREATE TABLE `restaurants` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `latitude` double NOT NULL,
  `longitude` double NOT NULL,
  `district_name` varchar(255) DEFAULT NULL,
  `cuisine_name` varchar(255) DEFAULT NULL,
  `rating` double DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=58247 DEFAULT CHARSET=latin1;

INSERT INTO restaurants(id, name, latitude, longitude, district_name, cuisine_name, rating)VALUES (12645, 'Polo Lounge', -6.2106388889, 106.8162222222, 'Tanah Abang', 'Western', 9.3333);
INSERT INTO restaurants(id, name, latitude, longitude, district_name, cuisine_name, rating) VALUES (38184, 'Martabak Orins', -6.20955, 106.8155833333, 'Tanah Abang', 'Desserts', 8.6);
INSERT INTO restaurants(id, name, latitude, longitude, district_name, cuisine_name, rating) VALUES (38861, 'Oto Bento', -6.2106722222, 106.8167416667, 'Tanah Abang', 'Japanese', 8.);
INSERT INTO restaurants(id, name, latitude, longitude, district_name, cuisine_name, rating) VALUES (38191, 'Skywalker Coffee', -6.2107111111, 106.8166666667, 'Tanah Abang', 'Coffee & Tea', 8.6829);
INSERT INTO restaurants(id, name, latitude, longitude, district_name, cuisine_name, rating) VALUES (12792, 'President Lounge', -6.2108, 106.8161694444, 'Tanah Abang', 'Western', 10);
INSERT INTO restaurants(id, name, latitude, longitude, district_name, cuisine_name, rating) VALUES (38803, 'Starbucks Coffee', -6.2108, 106.815985, 'Tanah Abang', 'Coffee & Tea', 9.2);
INSERT INTO restaurants(id, name, latitude, longitude, district_name, cuisine_name, rating) VALUES (47106, 'Waroeng Boenartie', -6.211001, 106.816228, 'Sudirman', 'Indonesian', 8.6);
INSERT INTO restaurants(id, name, latitude, longitude, district_name, cuisine_name, rating) VALUES (36606, 'Orenji Japanese Kitchen', -6.21106, 106.816326, 'Tanah Abang', 'Japanese', 8);
INSERT INTO restaurants(id, name, latitude, longitude, district_name, cuisine_name, rating) VALUES (58246, 'Bakso Benhil', -6.211078, 106.816369, 'Tanah Abang', 'Indonesian', 0);
INSERT INTO restaurants(id, name, latitude, longitude, district_name, cuisine_name, rating) VALUES (3098, 'Tjap Toean', -6.2089222222, 106.8180305556, 'Sudirman', 'Malaysian', 8);

INSERT INTO restaurants(id, name, latitude, longitude, district_name, cuisine_name, rating) VALUES (56400, 'Aroma Batavia Cafe & Resto', -6.176567, 106.841263, 'Senen', 'Indonesian', 0);
INSERT INTO restaurants(id, name, latitude, longitude, district_name, cuisine_name, rating) VALUES (36189, 'Waroeng Ikan Bakar Khas Papua', -6.176623, 106.841222, 'Senen', 'Indonesian', 0);
INSERT INTO restaurants(id, name, latitude, longitude, district_name, cuisine_name, rating) VALUES (36190, 'Es Pisang Ijo', -6.176623, 106.841222, 'Senen', 'Desserts', 0);
INSERT INTO restaurants(id, name, latitude, longitude, district_name, cuisine_name, rating) VALUES (36192, 'Makanan Khas Malang', -6.176623, 106.841222, 'Senen', 'Jawa', 0);
INSERT INTO restaurants(id, name, latitude, longitude, district_name, cuisine_name, rating) VALUES (36193, 'Nasi Goreng Semarang', -6.176623, 106.841222, 'Senen', 'Indonesian', 0);
INSERT INTO restaurants(id, name, latitude, longitude, district_name, cuisine_name, rating) VALUES (36194, 'Pindang Palembang Pondok Irmanti Sekayu', -6.176623, 106.841222, 'Senen', 'Palembang', 0);
INSERT INTO restaurants(id, name, latitude, longitude, district_name, cuisine_name, rating) VALUES (36195, 'Gado - Gado Boplo', -6.176623, 106.841222, 'Senen', 'Betawi', 0);
INSERT INTO restaurants(id, name, latitude, longitude, district_name, cuisine_name, rating) VALUES (31350, 'Kue Cubit Cimot', -6.176651, 106.841219, 'Senen', 'Desserts', 0);
INSERT INTO restaurants(id, name, latitude, longitude, district_name, cuisine_name, rating) VALUES (30989, 'Sushi Saki', -6.176652, 106.841219, 'Senen', 'Japanese', 0);
INSERT INTO restaurants(id, name, latitude, longitude, district_name, cuisine_name, rating) VALUES (30805, 'Steak 21', -6.176652, 106.841221, 'Senen', 'American', 0);

INSERT INTO restaurants(id, name, latitude, longitude, district_name, cuisine_name, rating) VALUES (10141, 'Gandasari', -6.194768, 106.81731, 'Thamrin', 'Sunda', 0);
INSERT INTO restaurants(id, name, latitude, longitude, district_name, cuisine_name, rating) VALUES (13003, 'Resto Ardhani', -6.194768, 106.81731, 'Thamrin', 'Indonesian', 0);
INSERT INTO restaurants(id, name, latitude, longitude, district_name, cuisine_name, rating) VALUES (13006, 'Resto Blueberry', -6.194768, 106.81731, 'Thamrin', 'Jawa', 0);
INSERT INTO restaurants(id, name, latitude, longitude, district_name, cuisine_name, rating) VALUES (13176, 'RM Maharani', -6.194768, 106.81731, 'Thamrin', 'Jawa', 0);
INSERT INTO restaurants(id, name, latitude, longitude, district_name, cuisine_name, rating) VALUES (13699, 'Sedap Murah', -6.194768, 106.81731, 'Thamrin', 'Jawa', 0);
INSERT INTO restaurants(id, name, latitude, longitude, district_name, cuisine_name, rating) VALUES (14102, 'Soto Betawi Ibu Annie', -6.194768, 106.81731, 'Thamrin', 'Betawi', 0);
INSERT INTO restaurants(id, name, latitude, longitude, district_name, cuisine_name, rating) VALUES (14142, 'Soto Kudus', -6.194768, 106.81731, 'Thamrin', 'Jawa', 0);
INSERT INTO restaurants(id, name, latitude, longitude, district_name, cuisine_name, rating) VALUES (24271, 'Burger King', -6.194768, 106.81731, 'Thamrin', 'Fast Food', 0);
INSERT INTO restaurants(id, name, latitude, longitude, district_name, cuisine_name, rating) VALUES (26444, 'Bakmi GM', -6.194768, 106.81731, 'Thamrin', 'Chinese', 0);
INSERT INTO restaurants(id, name, latitude, longitude, district_name, cuisine_name, rating) VALUES (28387, 'Kopi Oey', -6.194768, 106.81731, 'Thamrin', 'Coffee & Tea', 6);

CREATE TABLE `reservations` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `reservation_time` datetime DEFAULT NULL,
  `restaurant_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `restaurant_id` (`restaurant_id`),
  CONSTRAINT `reservations_ibfk_1` FOREIGN KEY (`restaurant_id`) REFERENCES `restaurants` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;