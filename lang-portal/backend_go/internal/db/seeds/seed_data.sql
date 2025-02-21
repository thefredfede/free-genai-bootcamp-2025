-- Insert words
INSERT INTO words (japanese, romaji, english, parts) VALUES
('猫', 'neko', 'cat', '["noun"]'),
('犬', 'inu', 'dog', '["noun"]'),
('鳥', 'tori', 'bird', '["noun"]'),
('魚', 'sakana', 'fish', '["noun"]'),
('車', 'kuruma', 'car', '["noun"]'),
('本', 'hon', 'book', '["noun"]'),
('学校', 'gakkou', 'school', '["noun"]'),
('先生', 'sensei', 'teacher', '["noun"]'),
('学生', 'gakusei', 'student', '["noun"]'),
('家', 'ie', 'house', '["noun"]'),
('食べる', 'taberu', 'eat', '["verb"]'),
('飲む', 'nomu', 'drink', '["verb"]'),
('走る', 'hashiru', 'run', '["verb"]'),
('読む', 'yomu', 'read', '["verb"]'),
('書く', 'kaku', 'write', '["verb"]'),
('見る', 'miru', 'see', '["verb"]'),
('聞く', 'kiku', 'listen', '["verb"]'),
('話す', 'hanasu', 'speak', '["verb"]'),
('買う', 'kau', 'buy', '["verb"]'),
('売る', 'uru', 'sell', '["verb"]');

-- Insert groups
INSERT INTO groups (name) VALUES
('Basic Vocabulary'),
('Animals'),
('Actions'),
('School'),
('Home');

-- Insert word_groups
INSERT INTO word_groups (word_id, group_id) VALUES
(1, 2), -- 猫 in Animals
(2, 2), -- 犬 in Animals
(3, 2), -- 鳥 in Animals
(4, 2), -- 魚 in Animals
(5, 1), -- 車 in Basic Vocabulary
(6, 1), -- 本 in Basic Vocabulary
(7, 4), -- 学校 in School
(8, 4), -- 先生 in School
(9, 4), -- 学生 in School
(10, 5), -- 家 in Home
(11, 3), -- 食べる in Actions
(12, 3), -- 飲む in Actions
(13, 3), -- 走る in Actions
(14, 3), -- 読む in Actions
(15, 3), -- 書く in Actions
(16, 3), -- 見る in Actions
(17, 3), -- 聞く in Actions
(18, 3), -- 話す in Actions
(19, 3), -- 買う in Actions
(20, 3); -- 売る in Actions

-- Insert study_sessions
INSERT INTO study_sessions (group_id, created_at, study_activity_id, activity_name, group_name) VALUES
(1, '2023-10-01T12:00:00Z', 1, 'Vocabulary Practice', 'Basic Vocabulary'),
(2, '2023-10-02T12:00:00Z', 2, 'Animal Names', 'Animals'),
(3, '2023-10-03T12:00:00Z', 3, 'Action Verbs', 'Actions'),
(4, '2023-10-04T12:00:00Z', 4, 'School Vocabulary', 'School'),
(5, '2023-10-05T12:00:00Z', 5, 'Home Vocabulary', 'Home');

-- Insert study_activities
INSERT INTO study_activities (study_session_id, group_id, created_at) VALUES
(1, 1, '2023-10-01T12:00:00Z'),
(2, 2, '2023-10-02T12:00:00Z'),
(3, 3, '2023-10-03T12:00:00Z'),
(4, 4, '2023-10-04T12:00:00Z'),
(5, 5, '2023-10-05T12:00:00Z');

-- Insert word_review_items
INSERT INTO word_review_items (word_id, study_session_id, correct, created_at) VALUES
(1, 1, true, '2023-10-01T12:05:00Z'),
(2, 1, false, '2023-10-01T12:10:00Z'),
(3, 2, true, '2023-10-02T12:05:00Z'),
(4, 2, true, '2023-10-02T12:10:00Z'),
(5, 3, false, '2023-10-03T12:05:00Z'),
(6, 3, true, '2023-10-03T12:10:00Z'),
(7, 4, true, '2023-10-04T12:05:00Z'),
(8, 4, false, '2023-10-04T12:10:00Z'),
(9, 5, true, '2023-10-05T12:05:00Z'),
(10, 5, true, '2023-10-05T12:10:00Z');