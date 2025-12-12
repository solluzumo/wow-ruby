-- =====================================
-- 1. Item
-- =====================================
INSERT INTO Item (id, name, price, required_level, max_stack, rarity, item_type)
VALUES
    ('11111111-1111-1111-1111-111111111111', 'Бронзовый Меч', 50, 1, 1, 'Обычный', 'Оружие'),
    ('22222222-2222-2222-2222-222222222222', 'Железный Топор', 120, 5, 1, 'Редкий', 'Оружие'),
    ('33333333-3333-3333-3333-333333333333', 'Кожаная Броня', 80, 3, 1, 'Обычный', 'Броня'),
    ('44444444-4444-4444-4444-444444444444', 'Стальной Щит', 200, 7, 1, 'Эпический', 'Броня'),
    ('55555555-5555-5555-5555-555555555555', 'Зелье Лечения', 20, 1, 10, 'Обычный', 'Расходуемое');


-- =====================================
-- 3. Equipment
-- =====================================
INSERT INTO Equipment (id, slot, capacity, durability)
VALUES
    ('11111111-1111-1111-1111-111111111111', 'Правая Рука', 100, 100),
    ('22222222-2222-2222-2222-222222222222', 'Правая Рука', 100, 100),
    ('33333333-3333-3333-3333-333333333333', 'Грудь', 50, 100),
    ('44444444-4444-4444-4444-444444444444', 'Левая Рука', 20, 100);

-- =====================================
-- 4. Weapon
-- =====================================
INSERT INTO Weapon (id, damage, speed, weapon_type)
VALUES
    ('11111111-1111-1111-1111-111111111111', '5-10', 2.0, 'Одноручный Меч'),
    ('22222222-2222-2222-2222-222222222222', '8-15', 2.5, 'Одноручный Топор');

-- =====================================
-- 5. Armor
-- =====================================
INSERT INTO Armor (id, armor_type, armor_value, set_name)
VALUES
    ('33333333-3333-3333-3333-333333333333', 'Легкая', 25, 'Кожаный Комплект'),
    ('44444444-4444-4444-4444-444444444444', 'Средняя', 40, 'Стальной Защитник');

-- =====================================
-- 6. QuestChain
-- =====================================
INSERT INTO QuestChain (id, chain_name)
VALUES ('cccccccc-cccc-cccc-cccc-cccccccccccc', 'Начало Героя');

-- =====================================
-- 7. Quest
-- =====================================
INSERT INTO Quest (id, name, reward_money, required_character_level, quest_level, is_repeatable, difficulty)
VALUES
    ('aaaa1111-1111-1111-1111-111111111111', 'Первая Кровь', 100, 1, 1, false, 1),
    ('aaaa2222-2222-2222-2222-222222222222', 'Защита Деревни', 300, 3, 4, false, 3);

-- =====================================
-- 8. ChainElement
-- =====================================
INSERT INTO ChainElement (quest_id, quest_chain_id, order_num)
VALUES
    ('aaaa1111-1111-1111-1111-111111111111', 'cccccccc-cccc-cccc-cccc-cccccccccccc', 1),
    ('aaaa2222-2222-2222-2222-222222222222', 'cccccccc-cccc-cccc-cccc-cccccccccccc', 2);

-- =====================================
-- 9. NPC
-- =====================================
INSERT INTO NPC (id, name, health, mana, level, tameable, faction, reaction, location, respawn_time)
VALUES
    ('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 'Страж Томас', 500, 100, 5, false, 'Альянс', 'Дружелюбный', 'Элвиннский Лес', 30),
    ('bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', 'Бандит', 300, 0, 3, false, 'Нейтральный', 'Враждебный', 'Элвиннский Лес', 45);

-- =====================================
-- 10. NPCItem
-- =====================================
INSERT INTO NPCItem (npc_id, item_id, quantity, drop_chance, source, quantity_limit)
VALUES
    ('bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', '11111111-1111-1111-1111-111111111111', 1, 10.5, 'Выпадение', NULL),
    ('bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', '55555555-5555-5555-5555-555555555555', 2, 30.0, 'Выпадение', NULL);

-- =====================================
-- 11. QuestItem
-- =====================================
INSERT INTO QuestItem (item_id, quest_id, item_role)
VALUES
    ('11111111-1111-1111-1111-111111111111', 'aaaa1111-1111-1111-1111-111111111111', 'used'),
    ('55555555-5555-5555-5555-555555555555', 'aaaa1111-1111-1111-1111-111111111111', 'reward');

-- =====================================
-- 12. CharacterQuest
-- =====================================
INSERT INTO CharacterQuest (npc_id, quest_id, npc_role)
VALUES
    ('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 'aaaa1111-1111-1111-1111-111111111111', 'Начинает'),
    ('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 'aaaa2222-2222-2222-2222-222222222222', 'Завершает');

-- =====================================
-- 13. Spell
-- =====================================
INSERT INTO Spell (
    spell_name, spell_level, description, magic_school, duration,
    mechanics, dispel_type, cooldown, cast_time, effect_radius,
    cost, required_level, range, player_usable, training_cost
)
VALUES
    ('Огненный Шар', 1, 'Бросает огненный шар, наносящий урон от огня.', 'Огонь', 10,
     'Снаряд', 'Магия', 8, 2.5, 30.0, 20, 1, 30, true, 100),

    ('Исцеление', 1, 'Восстанавливает здоровье союзнику.', 'Свет', 0,
     'Канал', 'Магия', 5, 2.0, 20.0, 15, 1, 25, true, 80);

-- =====================================
-- 14. Effect
-- =====================================
INSERT INTO Effect (id, spell_name, spell_level, effect_type, base_value, duration)
VALUES
    (gen_random_uuid(), 'Огненный Шар', 1, 'Урон', 50, 10),
    (gen_random_uuid(), 'Исцеление', 1, 'Лечение', 60, 0);

-- =====================================
-- 15. NPCSpell
-- =====================================
INSERT INTO NPCSpell (id, spell_name, spell_level, use_chance, interaction_type)
VALUES
    ('bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', 'Огненный Шар', 1, 25.0, 'Атаковать');

-- =====================================
-- 16. CharacterRole
-- =====================================
INSERT INTO CharacterRole (role_name, id)
VALUES
    ('Дающий Задания', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa'),
    ('Враг', 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb');
