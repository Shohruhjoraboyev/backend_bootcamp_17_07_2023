SELECT b.branches_name, COUNT(t.id) AS total_count 
FROM branches AS b 
JOIN branchprtransaction AS t ON b.id = t.branch_id
GROUP BY b.branches_name
ORDER BY total_count DESC;