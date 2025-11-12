CREATE TYPE rarity_enum AS ENUM ('Обычный', 'Редкий', 'Эпический', 'Легендарный');
CREATE TYPE source_type_enum AS ENUM ('Выпадение', 'Продажа');
CREATE TYPE armor_type_enum AS ENUM ('Легкая', 'Средняя', 'Тяжелая');
CREATE TYPE quest_repeatable_enum AS ENUM ('Да', 'Нет');
CREATE TYPE npc_role_enum AS ENUM ('Начинает', 'Завершает');
CREATE TYPE interaction_type_enum AS ENUM ('Атаковать', 'Тренировать');
CREATE TYPE item_type_enum AS ENUM ('Оружие', 'Броня', 'Расходуемое');
CREATE TYPE weapon_type_enum AS ENUM (
    'Кинжал',
    'Кастет',
    'Одноручный Топор',
    'Одноручное Дробящее Оружие',
    'Одноручный Меч',
    'Парные Клинки',
    'Древковое',
    'Посох',
    'Двуручный Топор',
    'Двуручное Дробящее Оружие',
    'Двуручный Меч',
    'Лук',
    'Арбалет',
    'Огнестрельное Оружие'
);

-- 1. Entity: Item
CREATE TABLE Item (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price INTEGER,
    required_level SMALLINT DEFAULT 1,
    max_stack SMALLINT NOT NULL,
    rarity rarity_enum NOT NULL,
    item_type item_type_enum NOT NULL,
    UNIQUE (name, item_type)
);

CREATE TABLE Users (
    id UUID PRIMARY KEY,
    login VARCHAR(100) NOT NULL,
    hash VARCHAR(100) NOT NULL
);

-- 20. Entity: Quest
CREATE TABLE Quest (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    reward_money INTEGER,
    required_character_level SMALLINT DEFAULT 1,
    quest_level SMALLINT NOT NULL,
    is_repeatable BOOLEAN NOT NULL,
    difficulty SMALLINT NOT NULL CHECK (difficulty BETWEEN 1 AND 5)
);

-- 21. Entity: QuestChain
CREATE TABLE QuestChain (
    id UUID PRIMARY KEY,
    chain_name VARCHAR(100) UNIQUE NOT NULL
);

-- 24. Entity: Spell
CREATE TABLE Spell (
    spell_name VARCHAR(100),
    spell_level SMALLINT CHECK (spell_level BETWEEN 1 AND 10),
    description TEXT NOT NULL,
    magic_school VARCHAR(50) NOT NULL,
    duration INTEGER,
    mechanics VARCHAR(100),
    dispel_type VARCHAR(50) NOT NULL,
    cooldown INTEGER,
    cast_time NUMERIC(3,2),
    effect_radius NUMERIC(4,2),
    cost INTEGER NOT NULL,
    required_level SMALLINT DEFAULT 1,
    range INTEGER,
    player_usable BOOLEAN NOT NULL,
    training_cost INTEGER,
    PRIMARY KEY (spell_name, spell_level)
);

-- 28. Entity: NPC
CREATE TABLE NPC (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    health INTEGER NOT NULL,
    mana INTEGER NOT NULL,
    level SMALLINT NOT NULL,
    tameable BOOLEAN NOT NULL,
    faction VARCHAR(50) NOT NULL,
    reaction VARCHAR(50) NOT NULL,
    location VARCHAR(100) NOT NULL,
    respawn_time INTEGER
);

-- 2. Entity: Schematic
CREATE TABLE Schematic (
    id UUID PRIMARY KEY,
    source_item INT NOT NULL REFERENCES Item(id),
    created_item INT NOT NULL REFERENCES Item(id),
    UNIQUE (source_item, created_item)
);

-- 3. Entity: Equipment Equipment.id должен совпадать с Item.id (одно-к-одному), поэтому INT PK REFERENCES Item(id)
CREATE TABLE Equipment (
    id INT PRIMARY KEY REFERENCES Item(id),
    slot VARCHAR(30) NOT NULL,
    capacity SMALLINT NOT NULL,
    durability SMALLINT DEFAULT 100
);

-- 5. Entity: QuestItem
CREATE TABLE QuestItem (
    item_id INT REFERENCES Item(id),
    quest_id INT REFERENCES Quest(id),
    item_role VARCHAR(20) NOT NULL CHECK (item_role IN ('used', 'reward')),
    PRIMARY KEY (item_id, quest_id)
);


-- 9. Entity: NPCItem
CREATE TABLE NPCItem (
    npc_id INT REFERENCES NPC(id),
    item_id INT REFERENCES Item(id),
    quantity SMALLINT,
    drop_chance NUMERIC(5,2) CHECK (drop_chance BETWEEN 0 AND 100),
    source source_type_enum NOT NULL,
    quantity_limit SMALLINT,
    PRIMARY KEY (npc_id, item_id)
);

-- 11. Entity: Armor
-- id должен совпадать с Equipment.id (1:1), поэтому INT PK REFERENCES Equipment(id)
CREATE TABLE Armor (
    id INT PRIMARY KEY REFERENCES Equipment(id),
    armor_type armor_type_enum NOT NULL,
    armor_value SMALLINT NOT NULL,
    set_name VARCHAR(100)
);

-- 12. Entity: Weapon
CREATE TABLE Weapon (
    id INT PRIMARY KEY REFERENCES Equipment(id),
    damage VARCHAR(20) NOT NULL,
    speed NUMERIC(3,2) NOT NULL,
    weapon_type weapon_type_enum  NOT NULL
);

-- 22. Entity: ChainElement
CREATE TABLE ChainElement (
    quest_id INT REFERENCES Quest(id),
    quest_chain_id INT REFERENCES QuestChain(id),
    order_num SMALLINT NOT NULL,
    PRIMARY KEY (quest_id, quest_chain_id)
);

-- 23. Entity: CharacterQuest
CREATE TABLE CharacterQuest (
    npc_id INT REFERENCES NPC(id),
    quest_id INT REFERENCES Quest(id),
    npc_role npc_role_enum NOT NULL,
    PRIMARY KEY (npc_id, quest_id)
);

-- 26. Entity: Effect
CREATE TABLE Effect (
    id UUID PRIMARY KEY,
    spell_name VARCHAR(100) NOT NULL,
    spell_level SMALLINT NOT NULL,
    effect_type VARCHAR(50) NOT NULL,
    base_value INTEGER NOT NULL,
    duration INTEGER,
    parent_id INT REFERENCES Effect(id),
    child_id INT REFERENCES Effect(id),
    UNIQUE (spell_name, spell_level, effect_type),
    FOREIGN KEY (spell_name, spell_level) REFERENCES Spell(spell_name, spell_level)
);

-- 27. Entity: NPCSpell
CREATE TABLE NPCSpell (
    id INT REFERENCES NPC(id),
    spell_name VARCHAR(100),
    spell_level SMALLINT,
    use_chance NUMERIC(5,2) CHECK (use_chance BETWEEN 0 AND 100),
    interaction_type interaction_type_enum NOT NULL,
    PRIMARY KEY (id, spell_name, spell_level),
    FOREIGN KEY (spell_name, spell_level) REFERENCES Spell(spell_name, spell_level)
);

-- 29. Entity: CharacterRole
CREATE TABLE CharacterRole (
    role_name VARCHAR(50),
    id INT REFERENCES NPC(id),
    PRIMARY KEY (role_name, id)
);

-- Indexes
CREATE INDEX idx_item_type ON Item(item_type);
CREATE INDEX idx_equipment_slot ON Equipment(slot);
CREATE INDEX idx_quest_level ON Quest(quest_level);
CREATE INDEX idx_npc_location ON NPC(location);
CREATE INDEX idx_spell_school ON Spell(magic_school);
CREATE INDEX idx_effect_type ON Effect(effect_type);
