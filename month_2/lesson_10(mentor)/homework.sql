-- HOMEWORK
-- miniProject
-- 1. Sale create/cancel qilinganda stafflarni balansiga pul hisoblanishini function/block ichida transaction bilan qiling

-- 2. agar staff bir kunda 10tadan ko'p va umumiy summasi 1 500 000 dan ko'p savdo qilsa 
-- 50 000 bonus berilishi kerak: balancega qo'shilishi va staff_transaction create qilinishi trigger bilan qiling

-- 3. pg_stat_activitydan username,query, exec_time ni exec time kamayish tartibida chiqaring
SELECT usename, query, clock_timestamp() - query_start AS exec_time
FROM pg_stat_activity
WHERE state = 'active'
ORDER BY exec_time DESC;

-- 4. pg_stat_user_indexesdan tablename, index_name, qancha ishlatilgani o'sish tartibida chiqaring
SELECT relname, indexrelname, idx_scan
FROM pg_stat_user_indexes
ORDER BY idx_scan DESC;

-- Themes:
-- https://www.postgresqltutorial.com:
-- 12,13,14 sections
-- Functions
-- Variables
-- Views