INSERT INTO public.status (status_name, status_active)
VALUES
('Recebido', true),
('Em Preparação', true),
('Pronto', true),
('Finalizado', true),
('Pendente Pagamento', true),
('Pago', true);

INSERT INTO public.category( category_name)
VALUES
('Lanche'),
('Acompanhamento'),
('Bebida'),
('Sobremesa');