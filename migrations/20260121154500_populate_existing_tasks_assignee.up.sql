-- Populate assignee_id for existing tasks
-- This migration fills in assignee_id based on reviewer_id for existing tasks
-- If you have a different logic, modify this query accordingly

-- Option 1: Set assignee_id to same as reviewer_id where both are NULL
UPDATE tasks 
SET assignee_id = reviewer_id 
WHERE assignee_id IS NULL AND reviewer_id IS NOT NULL;

-- Option 2: If you want to set a default assignee (e.g., user ID 1) for tasks without assignee
-- Uncomment the line below and modify the user ID as needed
-- UPDATE tasks SET assignee_id = 1 WHERE assignee_id IS NULL;
