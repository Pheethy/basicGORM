BEGIN;

INSERT INTO "roles" (
    "id",
    "title"
)
VALUES
    ('1', 'customer'),
    ('2', 'admin');

INSERT INTO "users" (
    "id",
    "username",
    "email",
    "password",
    "role_id"
)
VALUES
    ('1d6ca61e-df77-4b14-91da-e108c781797d', 'customer001', 'customer001@kawaii.com', '$2a$10$8KzaNdKIMyOkASCH4QvSKuEMIY7Jc3vcHDuSJvXLii1rvBNgz60a6', 1),
    ('47e9b29e-435d-43bc-8aeb-66d61dde0fa9', 'admin001', 'admin001@kawaii.com', '$2a$10$3qqNPE.TJpNGYCohjTgw9.v1z0ckovx95AmiEtUXcixGAgfW7.wCi', 2); --password = 123456


INSERT INTO "categories"
    (
        "title"
    )
VALUES
    ('food & beverage'),
    ('fashion'),
    ('gadget');

INSERT INTO "products"
    (
        "id",
        "title",
        "description",
        "price",
        "created_at",
        "updated_at"
    )
VALUES
    ('8e14ad55-48c1-491b-834e-6203edfc5936','Coffee', 'Just a food & beverage product', 150, '2023-03-10 00:03:59', '2023-03-10 00:03:59'),
    ('6b5cd101-7d72-4570-bda8-17eddd9a1547', 'Steak', 'Just a food & beverage product', 200, '2023-03-10 00:03:59', '2023-03-10 00:03:59'),
    ('ba3faeee-0a2b-44e7-81d7-f399bc342948', 'Shirt', 'Just a fashion product', 590, '2023-03-10 00:03:59', '2023-03-10 00:03:59'),
    ('f2506318-28e3-45a6-b0a0-9a1f486c2c10', 'Touser', 'Just a fashion product', 1490, '2023-03-10 00:03:59', '2023-03-10 00:03:59'),
    ('ce990149-ed3e-4921-8c3a-33d435170f7d', 'Phone', 'Just a gadget product', 33400, '2023-03-10 00:03:59', '2023-03-10 00:03:59'),
    ('ae36eb85-6c51-4085-8000-5797be838cca', 'Computer', 'Just a gadget product', 49000, '2023-03-10 00:03:59', '2023-03-10 00:03:59');

INSERT INTO "images"
    (
        "id",
        "filename",
        "url",
        "product_id",
        "created_at",
        "updated_at"
    )
VALUES
    ('c580fe73-afb3-47d1-a9df-eed24fdaea9b', 'fb1_1.jpg', 'https://i.pinimg.com/564x/4a/1c/4a/4a1c4a9755e4d3bdfcb45a1c3a58712f.jpg', '8e14ad55-48c1-491b-834e-6203edfc5936', '2023-03-10 00:03:59', '2023-03-10 00:03:59'),
    ('43bcd3fa-6f7f-4251-b196-f30ad4ea625e', 'fb1_2.jpg', 'https://i.pinimg.com/564x/4a/1c/4a/4a1c4a9755e4d3bdfcb45a1c3a58712f.jpg', '8e14ad55-48c1-491b-834e-6203edfc5936', '2023-03-10 00:03:59', '2023-03-10 00:03:59'),
    ('77d9e690-b722-4039-b0fe-5f7d9af0e6b4', 'fb1_3.jpg', 'https://i.pinimg.com/564x/4a/1c/4a/4a1c4a9755e4d3bdfcb45a1c3a58712f.jpg', '8e14ad55-48c1-491b-834e-6203edfc5936', '2023-03-10 00:03:59', '2023-03-10 00:03:59'),
    ('1d1eed38-3568-4e3e-9322-4c902b94c5b8', 'fb2_1.jpg', 'https://i.pinimg.com/564x/6d/ba/91/6dba91c1fdb5d4939c7e9d65420cbd4c.jpg', '6b5cd101-7d72-4570-bda8-17eddd9a1547', '2023-03-10 00:03:59', '2023-03-10 00:03:59'),
    ('f56c212a-16fd-4f8a-9091-03d2943c7f22', 'fb2_2.jpg', 'https://i.pinimg.com/564x/6d/ba/91/6dba91c1fdb5d4939c7e9d65420cbd4c.jpg', '6b5cd101-7d72-4570-bda8-17eddd9a1547', '2023-03-10 00:03:59', '2023-03-10 00:03:59'),
    ('6dfe9af7-1c48-4280-9805-60e7342ce2f7', 'fb2_3.jpg', 'https://i.pinimg.com/564x/6d/ba/91/6dba91c1fdb5d4939c7e9d65420cbd4c.jpg', '6b5cd101-7d72-4570-bda8-17eddd9a1547', '2023-03-10 00:03:59', '2023-03-10 00:03:59'),
    ('db2c59f0-434e-46b6-8184-e90c4bd15c3a', 'fs1_1.jpg', 'https://i.pinimg.com/564x/a0/6b/70/a06b708becbefa5d642392d7bf805429.jpg', 'ba3faeee-0a2b-44e7-81d7-f399bc342948', '2023-03-10 00:03:59', '2023-03-10 00:03:59'),
    ('4f1823d4-66e1-46de-bb15-8f56804bd810', 'fs1_2.jpg', 'https://i.pinimg.com/564x/a0/6b/70/a06b708becbefa5d642392d7bf805429.jpg', 'ba3faeee-0a2b-44e7-81d7-f399bc342948', '2023-03-10 00:03:59', '2023-03-10 00:03:59'),
    ('bdf45efe-6b87-4ae8-9695-9a356844494c', 'fs1_3.jpg', 'https://i.pinimg.com/564x/a0/6b/70/a06b708becbefa5d642392d7bf805429.jpg', 'ba3faeee-0a2b-44e7-81d7-f399bc342948', '2023-03-10 00:03:59', '2023-03-10 00:03:59'),
    ('251b8707-6a18-4cf9-b298-fec2a06586ca', 'fs2_1.jpg', 'https://i.pinimg.com/564x/e8/0a/0c/e80a0c4f562a942c01f6060a1e375a0b.jpg', 'f2506318-28e3-45a6-b0a0-9a1f486c2c10', '2023-03-10 00:03:59', '2023-03-10 00:03:59'),
    ('cadf3ebc-a1aa-4dc7-ab40-7e32d68ce4bc', 'fs2_2.jpg', 'https://i.pinimg.com/564x/e8/0a/0c/e80a0c4f562a942c01f6060a1e375a0b.jpg', 'f2506318-28e3-45a6-b0a0-9a1f486c2c10', '2023-03-10 00:03:59', '2023-03-10 00:03:59'),
    ('1e9bf281-76cf-4fc6-ba3b-22a66d9353b9', 'fs2_3.jpg', 'https://i.pinimg.com/564x/e8/0a/0c/e80a0c4f562a942c01f6060a1e375a0b.jpg', 'f2506318-28e3-45a6-b0a0-9a1f486c2c10', '2023-03-10 00:03:59', '2023-03-10 00:03:59'),
    ('e4c8ee7b-7c67-4d92-9955-d79f151bd40c', 'gt1_1.jpg', 'https://i.pinimg.com/564x/d5/95/e4/d595e4530aaa0fcdf4ff8e7bc17f4d86.jpg', 'ce990149-ed3e-4921-8c3a-33d435170f7d', '2023-03-10 00:03:59', '2023-03-10 00:03:59'),
    ('efae60af-94a5-4c2d-bb83-d3c5500c3c2e', 'gt1_2.jpg', 'https://i.pinimg.com/564x/d5/95/e4/d595e4530aaa0fcdf4ff8e7bc17f4d86.jpg', 'ce990149-ed3e-4921-8c3a-33d435170f7d', '2023-03-10 00:03:59', '2023-03-10 00:03:59'),
    ('1b4e1ec5-034a-441b-adcb-0da747ff49ef', 'gt1_3.jpg', 'https://i.pinimg.com/564x/d5/95/e4/d595e4530aaa0fcdf4ff8e7bc17f4d86.jpg', 'ce990149-ed3e-4921-8c3a-33d435170f7d', '2023-03-10 00:03:59', '2023-03-10 00:03:59'),
    ('df4912fc-c29b-48f1-a482-eaed6fb8f823', 'gt2_1.jpg', 'https://i.pinimg.com/564x/10/51/07/105107b2456059018b668f8d3e3989f6.jpg', 'ae36eb85-6c51-4085-8000-5797be838cca', '2023-03-10 00:03:59', '2023-03-10 00:03:59'),
    ('19d07a1f-342e-475d-8983-4a5ddc586ef1', 'gt2_2.jpg', 'https://i.pinimg.com/564x/10/51/07/105107b2456059018b668f8d3e3989f6.jpg', 'ae36eb85-6c51-4085-8000-5797be838cca', '2023-03-10 00:03:59', '2023-03-10 00:03:59'),
    ('dd65d3b2-3b50-49e3-9506-be66ef36810d', 'gt2_3.jpg', 'https://i.pinimg.com/564x/10/51/07/105107b2456059018b668f8d3e3989f6.jpg', 'ae36eb85-6c51-4085-8000-5797be838cca', '2023-03-10 00:03:59', '2023-03-10 00:03:59');

INSERT INTO "products_categories"
    (
        "product_id",
        "category_id"
    )
VALUES
    ('8e14ad55-48c1-491b-834e-6203edfc5936', 1),
    ('6b5cd101-7d72-4570-bda8-17eddd9a1547', 1),
    ('ba3faeee-0a2b-44e7-81d7-f399bc342948', 2),
    ('f2506318-28e3-45a6-b0a0-9a1f486c2c10', 2),
    ('ce990149-ed3e-4921-8c3a-33d435170f7d', 3),
    ('ae36eb85-6c51-4085-8000-5797be838cca', 3);

INSERT INTO "orders"
    (
        "id",
        "user_id",
        "contact",
        "address",
        "status"
    )
VALUES
    ('f3ea634f-b0d1-4261-bce3-58e63151d5f9', '1d6ca61e-df77-4b14-91da-e108c781797d', 'kawaii customer', '(330) 546-7713 5180 Richville Dr SW Navarre, Ohio(OH), 44662', 'COMPLETED'),
    ('18737b67-6b2e-45c8-8abc-2070f76fc816', '1d6ca61e-df77-4b14-91da-e108c781797d', 'kawaii customer', '(410) 256-8192 2260 Brimstone Pl Hanover, Maryland(MD), 21076', 'WAITING');

INSERT INTO "products_orders"
    (
        "order_id",
        "product_id",
        "qty"
    )
VALUES
    ('f3ea634f-b0d1-4261-bce3-58e63151d5f9','8e14ad55-48c1-491b-834e-6203edfc5936', 1),
    ('f3ea634f-b0d1-4261-bce3-58e63151d5f9','ae36eb85-6c51-4085-8000-5797be838cca', 2),
    ('18737b67-6b2e-45c8-8abc-2070f76fc816','8e14ad55-48c1-491b-834e-6203edfc5936', 1),
    ('18737b67-6b2e-45c8-8abc-2070f76fc816','ce990149-ed3e-4921-8c3a-33d435170f7d', 1);

INSERT INTO "transfer_slip" 
    (
        "id",
        "order_id",
        "filename",
        "url",
        "created_at"
    )
VALUES
    ('4bd7a0f5-c41f-4c1a-a997-0d965352fbb2', 'f3ea634f-b0d1-4261-bce3-58e63151d5f9', 'slip.jpg', 'https://i.pinimg.com/564x/a8/d4/f5/a8d4f5a620d22128c2b6d1a42c847560.jpg', '2023-03-01 23:21:00'),
    ('2e4b5a99-c2ec-488a-a4ed-e95306858838', '18737b67-6b2e-45c8-8abc-2070f76fc816', 'slip.jpg', 'https://i.pinimg.com/564x/a8/d4/f5/a8d4f5a620d22128c2b6d1a42c847560.jpg', '2023-03-01 23:21:00');

COMMIT;