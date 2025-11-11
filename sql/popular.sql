USE watermelon_db;

-- Limpa o banco antes de popular
DELETE FROM Usuario;
DELETE FROM Relatorio;
DELETE FROM Corpo_d_agua;
DELETE FROM Corpo_Localizacao;


INSERT INTO Usuario (Nome_de_usuario, E_mail, Senha) VALUES
  ('Erick Augusto', 'erick.augusto@aluno.ifsp.edu.br', 'password123'),
  ('Igor Angelo', 'igor@igor.com', 'l43gu30fl3g3'),
  ('Vinicius Gaioli', 'v@gaioli.br', 'all#mosso'),
  ('Vinicius Eduardo Murakami Moreira Pittoli', 'viniciuseduardomurakamimoreirapittoli@v.emmp.com.br', 'e4e5Bc4Nc6D');

INSERT INTO Corpo_d_agua (Nome, Tipo) VALUES
  ('Rio Jaguari-Mirim', 'rio'),
  ('Represa Billings', 'reservatorio'),
  ('Pampulha', 'lagoa'),
  ('Lagoa Rodrigo de Freitas', 'lagoa'),
  ('Baía de Todos os Santos', 'baia');

INSERT INTO Corpo_Localizacao (ID_Corpo_d_agua, ID_Localizacao) VALUES
  ((SELECT ID_Corpo_d_agua FROM Corpo_d_agua WHERE Nome = 'Rio Jaguari-Mirim'), (SELECT ID_Localizacao FROM Localizacao WHERE Cidade = 'Aguaí')),
  ((SELECT ID_Corpo_d_agua FROM Corpo_d_agua WHERE Nome = 'Rio Jaguari-Mirim'), (SELECT ID_Localizacao FROM Localizacao WHERE Cidade = 'São João da Boa Vista')),
  ((SELECT ID_Corpo_d_agua FROM Corpo_d_agua WHERE Nome = 'Represa Billings'), (SELECT ID_Localizacao FROM Localizacao WHERE Cidade = 'São Paulo')),
  ((SELECT ID_Corpo_d_agua FROM Corpo_d_agua WHERE Nome = 'Pampulha'), (SELECT ID_Localizacao FROM Localizacao WHERE Cidade = 'Belo Horizonte')),
  ((SELECT ID_Corpo_d_agua FROM Corpo_d_agua WHERE Nome = 'Lagoa Rodrigo de Freitas'), (SELECT ID_Localizacao FROM Localizacao WHERE Cidade = 'Rio de Janeiro')),
  ((SELECT ID_Corpo_d_agua FROM Corpo_d_agua WHERE Nome = 'Baía de Todos os Santos'), (SELECT ID_Localizacao FROM Localizacao WHERE Cidade = 'Salvador'));

-- Relatórios simples
INSERT INTO Relatorio (Tipo_de_relatorio, data, Descricao, Cor_aparente, ID_Autor, ID_Corpo_d_agua) VALUES
  ('simples', (SELECT CURRENT_DATE()), 'Água turva e marrom, indicando alta presença de sedimentos em suspensão, possível erosão nas margens e acúmulo de matéria orgânica. Pode conter poluentes agrícolas ou urbanos e apresentar baixo oxigênio dissolvido, prejudicando a fauna aquática.', 'marrom', (SELECT ID_usuario FROM Usuario WHERE Nome_de_usuario = 'Erick Augusto'), (SELECT ID_Corpo_d_agua FROM Corpo_d_agua WHERE Nome = 'Rio Jaguari-Mirim')),
  ('simples', (SELECT CURRENT_DATE()), 'Água esverdeada ou acastanhada, com turbidez moderada a alta, indicando presença de sedimentos, matéria orgânica e nutrientes em excesso. Possível eutrofização, com crescimento de algas e redução de oxigênio dissolvido, afetando peixes e outros organismos aquáticos. Contaminação por esgoto urbano e resíduos urbanos também é provável.', 'verde', (SELECT ID_usuario FROM Usuario WHERE Nome_de_usuario = 'Igor Angelo'), (SELECT ID_Corpo_d_agua FROM Corpo_d_agua WHERE Nome = 'Lagoa Rodrigo de Freitas')),
  ('simples', (SELECT CURRENT_DATE()), 'Água turva, com alta concentração de nutrientes, sedimentos e matéria orgânica. Indícios de eutrofização, crescimento de algas e baixa oxigenação. Poluição por esgoto doméstico, resíduos urbanos e drenagem urbana contribuem para a degradação da fauna e flora aquática.', 'turva', (SELECT ID_usuario FROM Usuario WHERE Nome_de_usuario = 'Vinicius Gaioli'), (SELECT ID_Corpo_d_agua FROM Corpo_d_agua WHERE Nome = 'Lagoa Rodrigo de Freitas'));

-- Relatórios avançados
INSERT INTO Relatorio (Tipo_de_relatorio, data, Descricao, Cor_aparente, ID_Autor, ID_Corpo_d_agua, Temperatura, Turbidez, Acidez, Odor, Sabor, Oxigenio_Dissolvido, Solidos_em_Suspensao) VALUES (
  'avancado',
  (SELECT CURRENT_DATE()),
  'Água escura, com turbidez elevada e presença de sedimentos e matéria orgânica. Possível eutrofização, com proliferação de algas e baixo oxigênio dissolvido. Poluição por esgoto, resíduos industriais e escoamento urbano compromete a qualidade e a vida aquática.',
  'azul',
  (SELECT ID_usuario FROM Usuario WHERE Nome_de_usuario = 'Erick Augusto'),
  (SELECT ID_Corpo_d_agua FROM Corpo_d_agua WHERE Nome = 'Represa Billings'),
  22.4,
  0.44,
  6.25,
  'cheiro de mato',
  'salgado',
  0.256,
  0.876
);

INSERT INTO Relatorio (Tipo_de_relatorio, data, Descricao, Cor_aparente, ID_Autor, ID_Corpo_d_agua, Temperatura, Turbidez, Acidez, Odor, Sabor, Oxigenio_Dissolvido, Solidos_em_Suspensao, Demanda_Bioquimica_de_Oxigenio, Nitrogenio_Total, Fosforo_Total, Metais_Pesados, Cloro_residual, Composto_Organico_Volatil, Solidos_Totais_Dissolvidos, Coliformes, Avaliacao_Biologica) VALUES (
  'avancado',
  (SELECT CURRENT_DATE()),
  'Água marrom, com turbidez variável devido a sedimentos em suspensão e correntes costeiras. Possível contaminação por esgoto doméstico, resíduos urbanos e portuários. Nutrientes em excesso podem favorecer crescimento de algas e reduzir oxigênio dissolvido, afetando fauna marinha.',
  'marrom',
  (SELECT ID_usuario FROM Usuario WHERE Nome_de_usuario = 'Vinicius Eduardo Murakami Moreira Pittoli'),
  (SELECT ID_Corpo_d_agua FROM Corpo_d_agua WHERE Nome = 'Baía de Todos os Santos'),
  22.4,
  0.44,
  6.25,
  'inodora',
  'terra',
  0.256,
  0.876,
  0.123,
  0.52,
  0.742,
  0.001,
  0.009,
  0.84,
  0.564,
  0.203,
  0.031
);

----------------------

