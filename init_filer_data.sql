--黎铭晖
INSERT INTO filer_account(filer_id, pay_id, token) VALUES (1, 1, 'vldYJ4tBM82rrwkT1fmiMotXUGhllWeZny3qOjkMvPhBePxyfzjg9dMWgdnxyAe5:1630928864');
--苏哲
INSERT INTO filer_account(filer_id, pay_id, token) VALUES (2, 2, 'vldYJptBM82rrwkT1fmiMotXUGhllWeZny3qOjkMvPhBePxyfzjg9dMWgdnxyAe5:1630928864');
--常春
INSERT INTO filer_account(filer_id, pay_id, token) VALUES (3, 3, 'vldYJ4tBM82rrpkT1fmiMotXUGhllWeZny3qOjkMvPhBePxyfzjg9dMWgdnxyAe5:1630928864');
--陈海洪
INSERT INTO filer_account(filer_id, pay_id, token) VALUES (4, 4, 'vldYJ4tBM82rrwkT1fmiMptXUGhllWeZny3qOjkMvPhBePxyfzjg9dMWgdnxyAe5:1630928864');
--黄虹兵
INSERT INTO filer_account(filer_id, pay_id, token) VALUES (5, 5, 'vldYJ4tBM82rrwkT1fmiMotXUGhllWeZnp3qOjkMvPhBePxyfzjg9dMWgdnxyAe5:1630928864');
--石志斌
INSERT INTO filer_account(filer_id, pay_id, token) VALUES (6, 6, 'vldYJ4tBM82rrwkT1fmiMotXUGhllWeZny3qOjkMpPhBePxyfzjg9dMWgdnxyAe5:1630928864');


-- pool
INSERT INTO filer_pool(node_id, node_name, location, mobile, email, create_time, update_time, is_valid) VALUES ('91fb8ea2-d435-4709-b933-1f7057b7f9ef', 'f01030435', '南沙', '13430750903', 'terilscaub@gmail.com',EXTRACT(epoch FROM to_timestamp('20210623111111', 'YYYYMMDDHHMISS')),EXTRACT(epoch FROM to_timestamp('20210623111111', 'YYYYMMDDHHMISS')), '0');

-- product -- 黎铭晖 920
INSERT INTO filer_product(product_id, product_name, node_id, cur_id, period, valid_plan, price, pledge_max, service_rate, node1, node2, shelve_time, create_time, update_time, product_state, is_valid)
VALUES ('91fb8ea2-d435-4709-b933-1f7057b7f9ef', '南沙920T-7', '91fb8ea2-d435-4709-b933-1f7057b7f9ef', 'FIL', '540', '3', '7200000000', '6625000000000', '0.12', 'no', 'no',EXTRACT(epoch FROM to_timestamp('20210630111111', 'YYYYMMDDHHMISS')), EXTRACT(epoch FROM now()::timestamp(0)),EXTRACT(epoch FROM now()::timestamp(0)), '0', '0');
-- product -- 苏哲 218
INSERT INTO filer_product(product_id, product_name, node_id, cur_id, period, valid_plan, price, pledge_max, service_rate, node1, node2, shelve_time, create_time, update_time, product_state, is_valid)
VALUES ('91fb8ea2-d435-4709-b933-1f7057b7f9eg', '南沙218T-7', '91fb8ea2-d435-4709-b933-1f7057b7f9ef', 'FIL', '540', '3', '7000000000', '1526000000000', '0.3', 'no', 'no',EXTRACT(epoch FROM to_timestamp('20210630111111', 'YYYYMMDDHHMISS')), EXTRACT(epoch FROM now()::timestamp(0)),EXTRACT(epoch FROM now()::timestamp(0)), '0', '0');
-- product -- 常春 200/陈海洪 40/黄虹兵 180
INSERT INTO filer_product(product_id, product_name, node_id, cur_id, period, valid_plan, price, pledge_max, service_rate, node1, node2, shelve_time, create_time, update_time, product_state, is_valid)
VALUES ('91fb8ea2-d435-4709-b933-1f7057b7f9eh', '南沙420T-10', '91fb8ea2-d435-4709-b933-1f7057b7f9ef', 'FIL', '540', '3', '5000000000', '2100000000000', '0.3', 'no', 'no',EXTRACT(epoch FROM to_timestamp('20210630111111', 'YYYYMMDDHHMISS')), EXTRACT(epoch FROM now()::timestamp(0)),EXTRACT(epoch FROM now()::timestamp(0)), '0', '0');
-- product -- 石志斌 721
INSERT INTO filer_product(product_id, product_name, node_id, cur_id, period, valid_plan, price, pledge_max, service_rate, node1, node2, shelve_time, create_time, update_time, product_state, is_valid)
VALUES ('91fb8ea2-d435-4709-b933-1f7057b7f9ej', '南沙721T-10', '91fb8ea2-d435-4709-b933-1f7057b7f9ef', 'FIL', '540', '3', '5729000000', '4131000000000', '0.3', 'no', 'no',EXTRACT(epoch FROM to_timestamp('20210630111111', 'YYYYMMDDHHMISS')), EXTRACT(epoch FROM now()::timestamp(0)),EXTRACT(epoch FROM now()::timestamp(0)), '0', '0');


-- -- product -- 黎铭晖 920
-- INSERT INTO public.filer_order
-- (order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, valid_time, end_time, order_state)
-- VALUES
-- ('order_id_01', '1', '0',
--  '91fb8ea2-d435-4709-b933-1f7057b7f9ef', -- 订单1 id
--  '920', --算力  920 T
--  '6625000000000',-- 6625 FIL
--  EXTRACT(epoch FROM to_timestamp('20210706111111', 'YYYYMMDDHHMISS')), --20210709 -3天 下单时间
--  EXTRACT(epoch FROM to_timestamp('20210706111111', 'YYYYMMDDHHMISS')), --修改时间
--  EXTRACT(epoch FROM to_timestamp('20210709111111', 'YYYYMMDDHHMISS')), --生效时间
--  EXTRACT(epoch FROM to_timestamp('20221231111111', 'YYYYMMDDHHMISS')), --2022-12-31  540天
--  '1'); --持仓状态 1 = 已生效
--
-- -- product -- 苏哲 218
-- INSERT INTO public.filer_order
-- (order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, valid_time, end_time, order_state)
-- VALUES
-- ('order_id_02', '2', '0',
--  '91fb8ea2-d435-4709-b933-1f7057b7f9eg',-- 订单2 id
--  '218', --算力  218 T
--  '1526000000000',-- 1526 FIL
--  EXTRACT(epoch FROM to_timestamp('20210628111111', 'YYYYMMDDHHMISS')), --20210701 -3天 下单时间
--  EXTRACT(epoch FROM to_timestamp('20210628111111', 'YYYYMMDDHHMISS')), --修改时间
--  EXTRACT(epoch FROM to_timestamp('20210701111111', 'YYYYMMDDHHMISS')), --生效时间
--  EXTRACT(epoch FROM to_timestamp('20221223111111', 'YYYYMMDDHHMISS')), -- 2022-12-23 540天
--  '1'); --持仓状态 1 = 已生效
--
-- -- product -- 常春 200
-- INSERT INTO public.filer_order
-- (order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, valid_time, end_time, order_state)
-- VALUES
-- ('order_id_03', '3', '0',
--  '91fb8ea2-d435-4709-b933-1f7057b7f9eh',-- 订单3 id
--  '200', --算力  200 T
--  '1000000000000',-- 1000 FIL
--  EXTRACT(epoch FROM to_timestamp('20210628111111', 'YYYYMMDDHHMISS')), --20210701 -3天 下单时间
--  EXTRACT(epoch FROM to_timestamp('20210628111111', 'YYYYMMDDHHMISS')), --修改时间
--  EXTRACT(epoch FROM to_timestamp('20210701111111', 'YYYYMMDDHHMISS')), --生效时间
--  EXTRACT(epoch FROM to_timestamp('20221223111111', 'YYYYMMDDHHMISS')), -- 2022-12-23  540天
--  '1'); --持仓状态 1 = 已生效
--
-- -- product 陈海洪 40
-- INSERT INTO public.filer_order
-- (order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, valid_time, end_time, order_state)
-- VALUES
-- ('order_id_04', '4', '0',
--  '91fb8ea2-d435-4709-b933-1f7057b7f9eh',-- 订单3 id
--  '40', --算力  40 T
--  '200000000000',-- 200 FIL
--  EXTRACT(epoch FROM to_timestamp('20210629111111', 'YYYYMMDDHHMISS')), --20210702 -3天 下单时间
--  EXTRACT(epoch FROM to_timestamp('20210629111111', 'YYYYMMDDHHMISS')), --修改时间
--  EXTRACT(epoch FROM to_timestamp('20210702111111', 'YYYYMMDDHHMISS')), --生效时间
--  EXTRACT(epoch FROM to_timestamp('20221224111111', 'YYYYMMDDHHMISS')), --2022-12-24 + 540天
--  '1'); --持仓状态 1 = 已生效
--
--
-- -- product -- 黎铭晖 920
-- INSERT INTO public.filer_order
-- (order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, valid_time, end_time, order_state)
-- VALUES
-- (gen_random_uuid(), '1', gen_random_uuid(),
--  '91fb8ea2-d435-4709-b933-1f7057b7f9ef', -- 订单1 id
--  '920', --算力  920 T
--  '6625000000000',-- 6625 FIL
--  EXTRACT(epoch FROM to_timestamp('20210706111111', 'YYYYMMDDHHMISS')), --20210709 -3天 下单时间
--  EXTRACT(epoch FROM to_timestamp('20210706111111', 'YYYYMMDDHHMISS')), --修改时间
--  EXTRACT(epoch FROM to_timestamp('20210709111111', 'YYYYMMDDHHMISS')), --生效时间
--  EXTRACT(epoch FROM to_timestamp('20221231111111', 'YYYYMMDDHHMISS')), --2022-12-31  540天
--  '1'); --持仓状态 1 = 已生效
--
-- -- product -- 苏哲 218
-- INSERT INTO public.filer_order
-- (order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, valid_time, end_time, order_state)
-- VALUES
-- (gen_random_uuid(), '2', gen_random_uuid(),
--  '91fb8ea2-d435-4709-b933-1f7057b7f9eg',-- 订单2 id
--  '218', --算力  218 T
--  '1526000000000',-- 1526 FIL
--  EXTRACT(epoch FROM to_timestamp('20210628111111', 'YYYYMMDDHHMISS')), --20210701 -3天 下单时间
--  EXTRACT(epoch FROM to_timestamp('20210628111111', 'YYYYMMDDHHMISS')), --修改时间
--  EXTRACT(epoch FROM to_timestamp('20210701111111', 'YYYYMMDDHHMISS')), --生效时间
--  EXTRACT(epoch FROM to_timestamp('20221223111111', 'YYYYMMDDHHMISS')), -- 2022-12-23 540天
--  '1'); --持仓状态 1 = 已生效
--
-- -- product -- 常春 200
-- INSERT INTO public.filer_order
-- (order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, valid_time, end_time, order_state)
-- VALUES
-- (gen_random_uuid(), '3', gen_random_uuid(),
--  '91fb8ea2-d435-4709-b933-1f7057b7f9eh',-- 订单3 id
--  '200', --算力  200 T
--  '1000000000000',-- 1000 FIL
--  EXTRACT(epoch FROM to_timestamp('20210628111111', 'YYYYMMDDHHMISS')), --20210701 -3天 下单时间
--  EXTRACT(epoch FROM to_timestamp('20210628111111', 'YYYYMMDDHHMISS')), --修改时间
--  EXTRACT(epoch FROM to_timestamp('20210701111111', 'YYYYMMDDHHMISS')), --生效时间
--  EXTRACT(epoch FROM to_timestamp('20221223111111', 'YYYYMMDDHHMISS')), -- 2022-12-23  540天
--  '1'); --持仓状态 1 = 已生效
--
-- -- product -- 常春 48
-- INSERT INTO public.filer_order
-- (order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, valid_time, end_time, order_state)
-- VALUES
-- (gen_random_uuid(), '3', gen_random_uuid(),
--  '91fb8ea2-d435-4709-b933-1f7057b7f9eh',-- 订单3 id
--  '48', --算力  200 T
--  '240000000000',-- 1000 FIL
--  EXTRACT(epoch FROM to_timestamp('20210816111111', 'YYYYMMDDHHMISS')), --20210701 -3天 下单时间
--  EXTRACT(epoch FROM to_timestamp('20210628111111', 'YYYYMMDDHHMISS')), --修改时间
--  EXTRACT(epoch FROM to_timestamp('20210819111111', 'YYYYMMDDHHMISS')), --生效时间
--  EXTRACT(epoch FROM to_timestamp('20230210111111', 'YYYYMMDDHHMISS')), --2023-02-10   540天
--  '1'); --持仓状态 1 = 已生效
--
-- -- product 陈海洪 40
-- INSERT INTO public.filer_order
-- (order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, valid_time, end_time, order_state)
-- VALUES
-- (gen_random_uuid(), '4', gen_random_uuid(),
--  '91fb8ea2-d435-4709-b933-1f7057b7f9eh',-- 订单3 id
--  '40', --算力  40 T
--  '200000000000',-- 200 FIL
--  EXTRACT(epoch FROM to_timestamp('20210629111111', 'YYYYMMDDHHMISS')), --20210702 -3天 下单时间
--  EXTRACT(epoch FROM to_timestamp('20210629111111', 'YYYYMMDDHHMISS')), --修改时间
--  EXTRACT(epoch FROM to_timestamp('20210702111111', 'YYYYMMDDHHMISS')), --生效时间
--  EXTRACT(epoch FROM to_timestamp('20221224111111', 'YYYYMMDDHHMISS')), --2022-12-24 + 540天
--  '1'); --持仓状态 1 = 已生效
--
--
--
-- -- product /黄虹兵 180
-- INSERT INTO public.filer_order
-- (order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, valid_time, end_time, order_state)
-- VALUES
-- (gen_random_uuid(), '5', gen_random_uuid(),
--  '91fb8ea2-d435-4709-b933-1f7057b7f9eh',-- 订单3 id
--  '180', --算力  180 T
--  '900000000000',-- 900 FIL
--  EXTRACT(epoch FROM to_timestamp('20210716111111', 'YYYYMMDDHHMISS')), --20210719 -3天 下单时间
--  EXTRACT(epoch FROM to_timestamp('20210716111111', 'YYYYMMDDHHMISS')), --修改时间
--  EXTRACT(epoch FROM to_timestamp('20210719111111', 'YYYYMMDDHHMISS')), --生效时间
--  EXTRACT(epoch FROM to_timestamp('20230110111111', 'YYYYMMDDHHMISS')), --2023-01-10  540天
--  '1'); --持仓状态 1 = 已生效
--
-- -- product /黄虹兵 320
-- INSERT INTO public.filer_order
-- (order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, valid_time, end_time, order_state)
-- VALUES
-- (gen_random_uuid(), '5', gen_random_uuid(),
--  '91fb8ea2-d435-4709-b933-1f7057b7f9eh',-- 订单3 id
--  '320', --算力  320T
--  '1600000000000',-- 1600 FIL
--  EXTRACT(epoch FROM to_timestamp('20210816111111', 'YYYYMMDDHHMISS')), --20210719 -3天 下单时间
--  EXTRACT(epoch FROM to_timestamp('20210816111111', 'YYYYMMDDHHMISS')), --修改时间
--  EXTRACT(epoch FROM to_timestamp('20210819111111', 'YYYYMMDDHHMISS')), --生效时间
--  EXTRACT(epoch FROM to_timestamp('20230210111111', 'YYYYMMDDHHMISS')), --2023-01-10  540天
--  '1'); --持仓状态 1 = 已生效
--
--
-- -- product /石志斌 200
-- INSERT INTO public.filer_order
-- (order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, valid_time, end_time, order_state)
-- VALUES
-- (gen_random_uuid(), '6', gen_random_uuid(),
--  '91fb8ea2-d435-4709-b933-1f7057b7f9ej',-- 订单4 id
--  '200', --算力  200 T
--  '1000000000000',-- 1000 FIL
--  EXTRACT(epoch FROM to_timestamp('20210720111111', 'YYYYMMDDHHMISS')), --20210723 -3天 下单时间
--  EXTRACT(epoch FROM to_timestamp('20210720111111', 'YYYYMMDDHHMISS')), --修改时间
--  EXTRACT(epoch FROM to_timestamp('20210723111111', 'YYYYMMDDHHMISS')), --生效时间
--  EXTRACT(epoch FROM to_timestamp('20230114111111', 'YYYYMMDDHHMISS')), --2023-01-14  540天
--  '1'); --持仓状态 1 = 已生效
--
-- -- product /石志斌 421
-- INSERT INTO public.filer_order
-- (order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, valid_time, end_time, order_state)
-- VALUES
-- (gen_random_uuid(), '6', gen_random_uuid(),
--  '91fb8ea2-d435-4709-b933-1f7057b7f9ej',-- 订单4 id
--  '421', --算力  421 T
--  '2531000000000',-- 2531 FIL
--  EXTRACT(epoch FROM to_timestamp('20210807111111', 'YYYYMMDDHHMISS')), --2021-08-10 -3天 下单时间
--  EXTRACT(epoch FROM to_timestamp('20210807111111', 'YYYYMMDDHHMISS')), --修改时间
--  EXTRACT(epoch FROM to_timestamp('20210810111111', 'YYYYMMDDHHMISS')), --生效时间
--  EXTRACT(epoch FROM to_timestamp('20230201111111', 'YYYYMMDDHHMISS')), --2023-02-01  540天
--  '1'); --持仓状态 1 = 已生效
--
-- -- product /石志斌 421
-- INSERT INTO public.filer_order
-- (order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, valid_time, end_time, order_state)
-- VALUES
-- (gen_random_uuid(), '6', gen_random_uuid(),
--  '91fb8ea2-d435-4709-b933-1f7057b7f9ej',-- 订单4 id
--  '100', --算力  100 T
--  '600000000000',-- 600 FIL
--  EXTRACT(epoch FROM to_timestamp('20210729111111', 'YYYYMMDDHHMISS')), --2021-08-01 -3天 下单时间
--  EXTRACT(epoch FROM to_timestamp('20210729111111', 'YYYYMMDDHHMISS')), --修改时间
--  EXTRACT(epoch FROM to_timestamp('20210801111111', 'YYYYMMDDHHMISS')), --生效时间
--  EXTRACT(epoch FROM to_timestamp('20230123111111', 'YYYYMMDDHHMISS')), --2023-01-23  540天
--  '1'); --持仓状态 1 = 已生效



 ---- 202108
-- product -- 黎铭晖 920
INSERT INTO public.filer_order
(order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, valid_time, end_time, order_state)
VALUES
(gen_random_uuid(), '1', gen_random_uuid(),
 '91fb8ea2-d435-4709-b933-1f7057b7f9ef', -- 订单1 id
 '920', --算力  920 T
 '6625000000000',-- 6625 FIL
 EXTRACT(epoch FROM to_timestamp('20210706111111', 'YYYYMMDDHHMISS')), --20210709 -3天 下单时间
 EXTRACT(epoch FROM to_timestamp('20210706111111', 'YYYYMMDDHHMISS')), --修改时间
 EXTRACT(epoch FROM to_timestamp('20210709111111', 'YYYYMMDDHHMISS')), --生效时间
 EXTRACT(epoch FROM to_timestamp('20221231111111', 'YYYYMMDDHHMISS')), --2022-12-31  540天
 '1'); --持仓状态 1 = 已生效

-- product -- 苏哲 218
INSERT INTO public.filer_order
(order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, valid_time, end_time, order_state)
VALUES
(gen_random_uuid(), '2', gen_random_uuid(),
 '91fb8ea2-d435-4709-b933-1f7057b7f9eg',-- 订单2 id
 '218', --算力  218 T
 '1526000000000',-- 1526 FIL
 EXTRACT(epoch FROM to_timestamp('20210628111111', 'YYYYMMDDHHMISS')), --20210701 -3天 下单时间
 EXTRACT(epoch FROM to_timestamp('20210628111111', 'YYYYMMDDHHMISS')), --修改时间
 EXTRACT(epoch FROM to_timestamp('20210701111111', 'YYYYMMDDHHMISS')), --生效时间
 EXTRACT(epoch FROM to_timestamp('20221223111111', 'YYYYMMDDHHMISS')), -- 2022-12-23 540天
 '1'); --持仓状态 1 = 已生效

-- product -- 常春 200
INSERT INTO public.filer_order
(order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, valid_time, end_time, order_state)
VALUES
(gen_random_uuid(), '3', gen_random_uuid(),
 '91fb8ea2-d435-4709-b933-1f7057b7f9eh',-- 订单3 id
 '200', --算力  200 T
 '1000000000000',-- 1000 FIL
 EXTRACT(epoch FROM to_timestamp('20210628111111', 'YYYYMMDDHHMISS')), --20210701 -3天 下单时间
 EXTRACT(epoch FROM to_timestamp('20210628111111', 'YYYYMMDDHHMISS')), --修改时间
 EXTRACT(epoch FROM to_timestamp('20210701111111', 'YYYYMMDDHHMISS')), --生效时间
 EXTRACT(epoch FROM to_timestamp('20221223111111', 'YYYYMMDDHHMISS')), -- 2022-12-23  540天
 '1'); --持仓状态 1 = 已生效

-- product -- 常春 48
INSERT INTO public.filer_order
(order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, valid_time, end_time, order_state)
VALUES
(gen_random_uuid(), '3', gen_random_uuid(),
 '91fb8ea2-d435-4709-b933-1f7057b7f9eh',-- 订单3 id
 '48', --算力  200 T
 '240000000000',-- 1000 FIL
 EXTRACT(epoch FROM to_timestamp('20210816111111', 'YYYYMMDDHHMISS')), --20210701 -3天 下单时间
 EXTRACT(epoch FROM to_timestamp('20210628111111', 'YYYYMMDDHHMISS')), --修改时间
 EXTRACT(epoch FROM to_timestamp('20210819111111', 'YYYYMMDDHHMISS')), --生效时间
 EXTRACT(epoch FROM to_timestamp('20230210111111', 'YYYYMMDDHHMISS')), --2023-02-10   540天
 '1'); --持仓状态 1 = 已生效

-- product 陈海洪 40
INSERT INTO public.filer_order
(order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, valid_time, end_time, order_state)
VALUES
(gen_random_uuid(), '4', gen_random_uuid(),
 '91fb8ea2-d435-4709-b933-1f7057b7f9eh',-- 订单3 id
 '40', --算力  40 T
 '200000000000',-- 200 FIL
 EXTRACT(epoch FROM to_timestamp('20210629111111', 'YYYYMMDDHHMISS')), --20210702 -3天 下单时间
 EXTRACT(epoch FROM to_timestamp('20210629111111', 'YYYYMMDDHHMISS')), --修改时间
 EXTRACT(epoch FROM to_timestamp('20210702111111', 'YYYYMMDDHHMISS')), --生效时间
 EXTRACT(epoch FROM to_timestamp('20221224111111', 'YYYYMMDDHHMISS')), --2022-12-24 + 540天
 '1'); --持仓状态 1 = 已生效

-- product /黄虹兵 180
INSERT INTO public.filer_order
(order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, valid_time, end_time, order_state)
VALUES
(gen_random_uuid(), '5', gen_random_uuid(),
 '91fb8ea2-d435-4709-b933-1f7057b7f9eh',-- 订单3 id
 '180', --算力  180 T
 '900000000000',-- 900 FIL
 EXTRACT(epoch FROM to_timestamp('20210716111111', 'YYYYMMDDHHMISS')), --20210719 -3天 下单时间
 EXTRACT(epoch FROM to_timestamp('20210716111111', 'YYYYMMDDHHMISS')), --修改时间
 EXTRACT(epoch FROM to_timestamp('20210719111111', 'YYYYMMDDHHMISS')), --生效时间
 EXTRACT(epoch FROM to_timestamp('20230110111111', 'YYYYMMDDHHMISS')), --2023-01-10  540天
 '1'); --持仓状态 1 = 已生效

-- product /黄虹兵 320
INSERT INTO public.filer_order
(order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, valid_time, end_time, order_state)
VALUES
(gen_random_uuid(), '5', gen_random_uuid(),
 '91fb8ea2-d435-4709-b933-1f7057b7f9eh',-- 订单3 id
 '320', --算力  320T
 '1600000000000',-- 1600 FIL
 EXTRACT(epoch FROM to_timestamp('20210816111111', 'YYYYMMDDHHMISS')), --20210719 -3天 下单时间
 EXTRACT(epoch FROM to_timestamp('20210816111111', 'YYYYMMDDHHMISS')), --修改时间
 EXTRACT(epoch FROM to_timestamp('20210819111111', 'YYYYMMDDHHMISS')), --生效时间
 EXTRACT(epoch FROM to_timestamp('20230210111111', 'YYYYMMDDHHMISS')), --2023-01-10  540天
 '1'); --持仓状态 1 = 已生效

-- product /石志斌 200
INSERT INTO public.filer_order
(order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, valid_time, end_time, order_state)
VALUES
(gen_random_uuid(), '6', gen_random_uuid(),
 '91fb8ea2-d435-4709-b933-1f7057b7f9ej',-- 订单4 id
 '200', --算力  200 T
 '1000000000000',-- 1000 FIL
 EXTRACT(epoch FROM to_timestamp('20210720111111', 'YYYYMMDDHHMISS')), --20210723 -3天 下单时间
 EXTRACT(epoch FROM to_timestamp('20210720111111', 'YYYYMMDDHHMISS')), --修改时间
 EXTRACT(epoch FROM to_timestamp('20210723111111', 'YYYYMMDDHHMISS')), --生效时间
 EXTRACT(epoch FROM to_timestamp('20230114111111', 'YYYYMMDDHHMISS')), --2023-01-14  540天
 '1'); --持仓状态 1 = 已生效

-- product /石志斌 421
INSERT INTO public.filer_order
(order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, valid_time, end_time, order_state)
VALUES
(gen_random_uuid(), '6', gen_random_uuid(),
 '91fb8ea2-d435-4709-b933-1f7057b7f9ej',-- 订单4 id
 '421', --算力  421 T
 '2531000000000',-- 2531 FIL
 EXTRACT(epoch FROM to_timestamp('20210807111111', 'YYYYMMDDHHMISS')), --2021-08-10 -3天 下单时间
 EXTRACT(epoch FROM to_timestamp('20210807111111', 'YYYYMMDDHHMISS')), --修改时间
 EXTRACT(epoch FROM to_timestamp('20210810111111', 'YYYYMMDDHHMISS')), --生效时间
 EXTRACT(epoch FROM to_timestamp('20230201111111', 'YYYYMMDDHHMISS')), --2023-02-01  540天
 '1'); --持仓状态 1 = 已生效

-- product /石志斌 421
INSERT INTO public.filer_order
(order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, valid_time, end_time, order_state)
VALUES
(gen_random_uuid(), '6', gen_random_uuid(),
 '91fb8ea2-d435-4709-b933-1f7057b7f9ej',-- 订单4 id
 '100', --算力  100 T
 '600000000000',-- 600 FIL
 EXTRACT(epoch FROM to_timestamp('20210729111111', 'YYYYMMDDHHMISS')), --2021-08-01 -3天 下单时间
 EXTRACT(epoch FROM to_timestamp('20210729111111', 'YYYYMMDDHHMISS')), --修改时间
 EXTRACT(epoch FROM to_timestamp('20210801111111', 'YYYYMMDDHHMISS')), --生效时间
 EXTRACT(epoch FROM to_timestamp('20230123111111', 'YYYYMMDDHHMISS')), --2023-01-23  540天
 '1'); --持仓状态 1 = 已生效

-- INSERT INTO public.filer_order (order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, end_time, order_state) VALUES ('5148b262-ed24-45ac-851d-f7d231d6f2c4', '6', 'c0fcc02c-1f35-46f9-bedf-ddf7a7c9dae5', '91fb8ea2-d435-4709-b933-1f7057b7f9ej', '100', '600000000000', '1627557071', '1631106703', '1674472271', '1');
-- INSERT INTO public.filer_order (order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, end_time, order_state) VALUES ('080f7c0c-259b-41eb-9b07-ce4b8ea314fe', '6', '58bb9670-5359-4a9f-af77-34a51fab79f7', '91fb8ea2-d435-4709-b933-1f7057b7f9ej', '421', '2531000000000', '1628334671', '1631106704', '1675249871', '1');
-- INSERT INTO public.filer_order (order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, end_time, order_state) VALUES ('6df1b4b1-f364-4cb3-b0e4-cacec80174bf', '2', '06a6d883-c279-4e0c-8070-21084797ed9e', '91fb8ea2-d435-4709-b933-1f7057b7f9eg', '218', '1526000000000', '1624878671', '1631106700', '1671793871', '1');
-- INSERT INTO public.filer_order (order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, end_time, order_state) VALUES ('c7d1a962-0da3-4160-850a-4e03aacc4eee', '3', '0c7e3a9c-c3f5-4b4e-baa1-d72a83bfb1ff', '91fb8ea2-d435-4709-b933-1f7057b7f9eh', '200', '1000000000000', '1624878671', '1631106700', '1671793871', '1');
-- INSERT INTO public.filer_order (order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, end_time, order_state) VALUES ('99c2ecbd-e92f-4db3-96a8-76ed47e30bea', '4', 'bc3c6131-96c0-4c73-a15c-1c941b97ceaf', '91fb8ea2-d435-4709-b933-1f7057b7f9eh', '40', '200000000000', '1624965071', '1631106700', '1671880271', '1');
-- INSERT INTO public.filer_order (order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, end_time, order_state) VALUES ('ad9b3afc-ed4d-4ca0-b008-fabf937961bf', '1', '9c36abdb-e406-4d60-bb19-2c75efc53cd6', '91fb8ea2-d435-4709-b933-1f7057b7f9ef', '920', '6625000000000', '1625569871', '1631106700', '1672485071', '1');
-- INSERT INTO public.filer_order (order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, end_time, order_state) VALUES ('ea2ecf10-a05f-4c35-984c-4546c15e73b4', '5', 'cf7da500-2d2b-414b-92dd-be2e8e07dbaa', '91fb8ea2-d435-4709-b933-1f7057b7f9eh', '180', '900000000000', '1626433871', '1631106701', '1673349071', '1');
-- INSERT INTO public.filer_order (order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, end_time, order_state) VALUES ('2bb041b5-32dd-4187-ac10-2cd4f6479715', '6', 'f019a909-1df8-4ab2-8be0-c2ca0c533061', '91fb8ea2-d435-4709-b933-1f7057b7f9ej', '200', '1000000000000', '1626779471', '1631106702', '1673694671', '1');