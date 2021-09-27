-- pool
INSERT INTO filer_pool(node_id, node_name, location, mobile, email, create_time, update_time, is_valid)
VALUES ('e70e8a2e-6575-40ff-a4d5-62b114972c66', 'f01259449', '502POOL', '13430750903', 'terilscaub@gmail.com',
        EXTRACT(epoch FROM to_timestamp('20210920111111', 'YYYYMMDDHHMISS')),
        EXTRACT(epoch FROM to_timestamp('20210920111111', 'YYYYMMDDHHMISS')), '0');

INSERT INTO filer_product(product_id, product_name, node_id, cur_id, period, valid_plan, price, pledge_max,
                          service_rate, note1, note2, shelve_time, create_time, update_time, product_state, is_valid)
VALUES ('91fb8ea2-d435-4709-b933-1f7057b7f9ef', '502矿池1期', 'e70e8a2e-6575-40ff-a4d5-62b114972c66', 'FIL', '540', '1',
        '5150000000', '6625000000000', '0.3', 'no', 'no',
        EXTRACT(epoch FROM to_timestamp('20210920111111', 'YYYYMMDDHHMISS')), EXTRACT(epoch FROM now()::timestamp(0)),
        EXTRACT(epoch FROM now()::timestamp(0)), '0', '0');

INSERT INTO filer_account_info(filer_id, filer_name, reg_time, mobile, email, is_valid)
VALUES ('81882469-1840-46fc-ae37-7d252c885193', 'Dylan',
        EXTRACT(epoch FROM to_timestamp('20210927111111', 'YYYYMMDDHHMISS')), '', '', '0');

INSERT INTO filer_account_info(filer_id, filer_name, reg_time, mobile, email, is_valid)
VALUES ('fc98db17-6911-47a3-b58a-30083c115004', 'ZBC',
        EXTRACT(epoch FROM to_timestamp('20210927111111', 'YYYYMMDDHHMISS')), '', '', '0');

INSERT INTO filer_account_info(filer_id, filer_name, reg_time, mobile, email, is_valid)
VALUES ('ba9d2f54-76b6-47b3-ac20-670f70bd4715', 'DaYang',
        EXTRACT(epoch FROM to_timestamp('20210927111111', 'YYYYMMDDHHMISS')), '', '', '0');

INSERT INTO filer_account_info(filer_id, filer_name, reg_time, mobile, email, is_valid)
VALUES ('5c11e46d-884b-4922-ab20-c33d015aa62d', 'Alienegra',
        EXTRACT(epoch FROM to_timestamp('20210927111111', 'YYYYMMDDHHMISS')), '', '', '0');

INSERT INTO filer_account_info(filer_id, filer_name, reg_time, mobile, email, is_valid)
VALUES ('fdae17dd-a71c-4c9d-9d15-0ce6841a938c', 'Vincent',
        EXTRACT(epoch FROM to_timestamp('20210927111111', 'YYYYMMDDHHMISS')), '', '', '0');

INSERT INTO filer_account_info(filer_id, filer_name, reg_time, mobile, email, is_valid)
VALUES ('98226dc7-73cb-421f-bf16-c2bcf58d8f6c', 'Terrill',
        EXTRACT(epoch FROM to_timestamp('20210927111111', 'YYYYMMDDHHMISS')), '', '', '0');

INSERT INTO filer_account_info(filer_id, filer_name, reg_time, mobile, email, is_valid)
VALUES ('87c380e1-a471-4cab-a176-19966e108ded', 'XP',
        EXTRACT(epoch FROM to_timestamp('20210927111111', 'YYYYMMDDHHMISS')), '', '', '0');

INSERT INTO filer_account_info(filer_id, filer_name, reg_time, mobile, email, is_valid)
VALUES ('956ba5c4-42f6-4889-8944-c02952a9c01c', 'JC',
        EXTRACT(epoch FROM to_timestamp('20210927111111', 'YYYYMMDDHHMISS')), '', '', '0');

INSERT INTO filer_account_info(filer_id, filer_name, reg_time, mobile, email, is_valid)
VALUES ('43fb37fb-1c42-45a4-b651-575ac782f4ee', 'JianBao',
        EXTRACT(epoch FROM to_timestamp('20210927111111', 'YYYYMMDDHHMISS')), '', '', '0');

INSERT INTO filer_account_info(filer_id, filer_name, reg_time, mobile, email, is_valid)
VALUES ('482f0608-db00-4a63-82d1-8df40688aef4', 'Grace',
        EXTRACT(epoch FROM to_timestamp('20210927111111', 'YYYYMMDDHHMISS')), '', '', '0');

INSERT INTO filer_account_info(filer_id, filer_name, reg_time, mobile, email, is_valid)
VALUES ('50c47f49-8851-4fc2-9754-8894303d268d', 'SUM',
        EXTRACT(epoch FROM to_timestamp('20210927111111', 'YYYYMMDDHHMISS')), '', '', '0');

-- 81882469-1840-46fc-ae37-7d252c885193,Dylan       Dylan	    440	    85.47
INSERT INTO filer_order
(order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, valid_time, end_time,
 order_state)
VALUES (gen_random_uuid(), '81882469-1840-46fc-ae37-7d252c885193', gen_random_uuid(),
        '91fb8ea2-d435-4709-b933-1f7057b7f9ef',
        '85.47', --算力  920 T
        '440000000000',-- 6625 FIL
        EXTRACT(epoch FROM to_timestamp('20210925111111', 'YYYYMMDDHHMISS')), --20210709 -3天 下单时间
        EXTRACT(epoch FROM to_timestamp('20210925111111', 'YYYYMMDDHHMISS')), --修改时间
        EXTRACT(epoch FROM to_timestamp('20210926111111', 'YYYYMMDDHHMISS')), --生效时间
        EXTRACT(epoch FROM to_timestamp('20230319111111', 'YYYYMMDDHHMISS')), --2023-03-19  540天
        '1');

-- fc98db17-6911-47a3-b58a-30083c115004,ZBC         ZBC	        560	    108.75
INSERT INTO filer_order
(order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, valid_time, end_time,
 order_state)
VALUES (gen_random_uuid(), 'fc98db17-6911-47a3-b58a-30083c115004', gen_random_uuid(),
        '91fb8ea2-d435-4709-b933-1f7057b7f9ef',
        '108.75', --算力  920 T
        '560000000000',-- 6625 FIL
        EXTRACT(epoch FROM to_timestamp('20210925111111', 'YYYYMMDDHHMISS')), --20210709 -3天 下单时间
        EXTRACT(epoch FROM to_timestamp('20210925111111', 'YYYYMMDDHHMISS')), --修改时间
        EXTRACT(epoch FROM to_timestamp('20210926111111', 'YYYYMMDDHHMISS')), --生效时间
        EXTRACT(epoch FROM to_timestamp('20230319111111', 'YYYYMMDDHHMISS')), --2023-03-19  540天
        '1');

-- ba9d2f54-76b6-47b3-ac20-670f70bd4715,DaYang      DaYang	    60	    11.65
INSERT INTO filer_order
(order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, valid_time, end_time,
 order_state)
VALUES (gen_random_uuid(), 'ba9d2f54-76b6-47b3-ac20-670f70bd4715', gen_random_uuid(),
        '91fb8ea2-d435-4709-b933-1f7057b7f9ef',
        '11.65', --算力  920 T
        '60000000000',-- 6625 FIL
        EXTRACT(epoch FROM to_timestamp('20210925111111', 'YYYYMMDDHHMISS')), --20210709 -3天 下单时间
        EXTRACT(epoch FROM to_timestamp('20210925111111', 'YYYYMMDDHHMISS')), --修改时间
        EXTRACT(epoch FROM to_timestamp('20210926111111', 'YYYYMMDDHHMISS')), --生效时间
        EXTRACT(epoch FROM to_timestamp('20230319111111', 'YYYYMMDDHHMISS')), --2023-03-19  540天
        '1');
-- 5c11e46d-884b-4922-ab20-c33d015aa62d,Alienegra   Alienegra	180	    34.86
INSERT INTO filer_order
(order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, valid_time, end_time,
 order_state)
VALUES (gen_random_uuid(), '5c11e46d-884b-4922-ab20-c33d015aa62d', gen_random_uuid(),
        '91fb8ea2-d435-4709-b933-1f7057b7f9ef',
        '34.86', --算力  920 T
        '180000000000',-- 6625 FIL
        EXTRACT(epoch FROM to_timestamp('20210925111111', 'YYYYMMDDHHMISS')), --20210709 -3天 下单时间
        EXTRACT(epoch FROM to_timestamp('20210925111111', 'YYYYMMDDHHMISS')), --修改时间
        EXTRACT(epoch FROM to_timestamp('20210926111111', 'YYYYMMDDHHMISS')), --生效时间
        EXTRACT(epoch FROM to_timestamp('20230319111111', 'YYYYMMDDHHMISS')), --2023-03-19  540天
        '1');

-- fdae17dd-a71c-4c9d-9d15-0ce6841a938c,Vincent       Vincent	    55.45	10.77
INSERT INTO filer_order
(order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, valid_time, end_time,
 order_state)
VALUES (gen_random_uuid(), 'fdae17dd-a71c-4c9d-9d15-0ce6841a938c', gen_random_uuid(),
        '91fb8ea2-d435-4709-b933-1f7057b7f9ef',
        '10.77', --算力  920 T
        '55450000000',-- 6625 FIL
        EXTRACT(epoch FROM to_timestamp('20210925111111', 'YYYYMMDDHHMISS')), --20210709 -3天 下单时间
        EXTRACT(epoch FROM to_timestamp('20210925111111', 'YYYYMMDDHHMISS')), --修改时间
        EXTRACT(epoch FROM to_timestamp('20210926111111', 'YYYYMMDDHHMISS')), --生效时间
        EXTRACT(epoch FROM to_timestamp('20230319111111', 'YYYYMMDDHHMISS')), --2023-03-19  540天
        '1');

-- 98226dc7-73cb-421f-bf16-c2bcf58d8f6c,Terrill     Terrill	    568	    110.5
INSERT INTO filer_order
(order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, valid_time, end_time,
 order_state)
VALUES (gen_random_uuid(), '98226dc7-73cb-421f-bf16-c2bcf58d8f6c', gen_random_uuid(),
        '91fb8ea2-d435-4709-b933-1f7057b7f9ef',
        '110.5', --算力  920 T
        '568000000000',-- 6625 FIL
        EXTRACT(epoch FROM to_timestamp('20210925111111', 'YYYYMMDDHHMISS')), --20210709 -3天 下单时间
        EXTRACT(epoch FROM to_timestamp('20210925111111', 'YYYYMMDDHHMISS')), --修改时间
        EXTRACT(epoch FROM to_timestamp('20210926111111', 'YYYYMMDDHHMISS')), --生效时间
        EXTRACT(epoch FROM to_timestamp('20230319111111', 'YYYYMMDDHHMISS')), --2023-03-19  540天
        '1');
-- 87c380e1-a471-4cab-a176-19966e108ded,XP          XP	        60	    11.65
INSERT INTO filer_order
(order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, valid_time, end_time,
 order_state)
VALUES (gen_random_uuid(), '87c380e1-a471-4cab-a176-19966e108ded', gen_random_uuid(),
        '91fb8ea2-d435-4709-b933-1f7057b7f9ef',
        '11.65', --算力  920 T
        '60000000000',-- 6625 FIL
        EXTRACT(epoch FROM to_timestamp('20210925111111', 'YYYYMMDDHHMISS')), --20210709 -3天 下单时间
        EXTRACT(epoch FROM to_timestamp('20210925111111', 'YYYYMMDDHHMISS')), --修改时间
        EXTRACT(epoch FROM to_timestamp('20210926111111', 'YYYYMMDDHHMISS')), --生效时间
        EXTRACT(epoch FROM to_timestamp('20230319111111', 'YYYYMMDDHHMISS')), --2023-03-19  540天
        '1');
-- 956ba5c4-42f6-4889-8944-c02952a9c01c,JC          JC	        6	    1.16
INSERT INTO filer_order
(order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, valid_time, end_time,
 order_state)
VALUES (gen_random_uuid(), '956ba5c4-42f6-4889-8944-c02952a9c01c', gen_random_uuid(),
        '91fb8ea2-d435-4709-b933-1f7057b7f9ef',
        '1.16', --算力  920 T
        '6000000000',-- 6625 FIL
        EXTRACT(epoch FROM to_timestamp('20210925111111', 'YYYYMMDDHHMISS')), --20210709 -3天 下单时间
        EXTRACT(epoch FROM to_timestamp('20210925111111', 'YYYYMMDDHHMISS')), --修改时间
        EXTRACT(epoch FROM to_timestamp('20210926111111', 'YYYYMMDDHHMISS')), --生效时间
        EXTRACT(epoch FROM to_timestamp('20230319111111', 'YYYYMMDDHHMISS')), --2023-03-19  540天
        '1');
-- 43fb37fb-1c42-45a4-b651-575ac782f4ee,JianBao     JianBao	    60	    11.65
INSERT INTO filer_order
(order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, valid_time, end_time,
 order_state)
VALUES (gen_random_uuid(), '43fb37fb-1c42-45a4-b651-575ac782f4ee', gen_random_uuid(),
        '91fb8ea2-d435-4709-b933-1f7057b7f9ef',
        '11.65', --算力  920 T
        '60000000000',-- 6625 FIL
        EXTRACT(epoch FROM to_timestamp('20210925111111', 'YYYYMMDDHHMISS')), --20210709 -3天 下单时间
        EXTRACT(epoch FROM to_timestamp('20210925111111', 'YYYYMMDDHHMISS')), --修改时间
        EXTRACT(epoch FROM to_timestamp('20210926111111', 'YYYYMMDDHHMISS')), --生效时间
        EXTRACT(epoch FROM to_timestamp('20230319111111', 'YYYYMMDDHHMISS')), --2023-03-19  540天
        '1');
-- 482f0608-db00-4a63-82d1-8df40688aef4,Grace       Grace	    198.33	38.44
INSERT INTO filer_order
(order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, valid_time, end_time,
 order_state)
VALUES (gen_random_uuid(), '482f0608-db00-4a63-82d1-8df40688aef4', gen_random_uuid(),
        '91fb8ea2-d435-4709-b933-1f7057b7f9ef',
        '38.44', --算力  920 T
        '198330000000',-- 6625 FIL
        EXTRACT(epoch FROM to_timestamp('20210925111111', 'YYYYMMDDHHMISS')), --20210709 -3天 下单时间
        EXTRACT(epoch FROM to_timestamp('20210925111111', 'YYYYMMDDHHMISS')), --修改时间
        EXTRACT(epoch FROM to_timestamp('20210926111111', 'YYYYMMDDHHMISS')), --生效时间
        EXTRACT(epoch FROM to_timestamp('20230319111111', 'YYYYMMDDHHMISS')), --2023-03-19  540天
        '1');
-- 50c47f49-8851-4fc2-9754-8894303d268d,SUM         SUM	        2187.78	425
INSERT INTO filer_order
(order_id, filer_id, pay_flow, product_id, hold_power, pay_amount, order_time, update_time, valid_time, end_time,
 order_state)
VALUES (gen_random_uuid(), '50c47f49-8851-4fc2-9754-8894303d268d', gen_random_uuid(),
        '91fb8ea2-d435-4709-b933-1f7057b7f9ef',
        '425', --算力  920 T
        '2187780000000',-- 6625 FIL
        EXTRACT(epoch FROM to_timestamp('20210925111111', 'YYYYMMDDHHMISS')), --20210709 -3天 下单时间
        EXTRACT(epoch FROM to_timestamp('20210925111111', 'YYYYMMDDHHMISS')), --修改时间
        EXTRACT(epoch FROM to_timestamp('20210926111111', 'YYYYMMDDHHMISS')), --生效时间
        EXTRACT(epoch FROM to_timestamp('20230319111111', 'YYYYMMDDHHMISS')), --2023-03-19  540天
        '1');