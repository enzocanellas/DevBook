insert into usuarios(nome, nick, email, senha)
values
("Usuário 1", "usuario_1", "usuario1@gmail.com", " $2a$10$2MupedtBOsTIoGEVHFuQ1emPoLBeHg7uTqoQnnWxn0.GEXen9A3v2"),
("Usuário 2", "usuario_2", "usuario2@gmail.com", " $2a$10$2MupedtBOsTIoGEVHFuQ1emPoLBeHg7uTqoQnnWxn0.GEXen9A3v2"),
("Usuário 3", "usuario_3", "usuario3@gmail.com", " $2a$10$2MupedtBOsTIoGEVHFuQ1emPoLBeHg7uTqoQnnWxn0.GEXen9A3v2");

insert into seguidores(usuario_id, seguidor_id)
values
(1, 2), 
(3, 1),
(1, 3);