package dto

import (
	"github.com/ruziba3vich/sahiy_management/internal/domain/section"
	domain "github.com/ruziba3vich/sahiy_management/internal/domain/task"
	"github.com/ruziba3vich/sahiy_management/internal/domain/taskstatus"
)

type CreateTaskRequest struct {
	ParentID    *int64 `json:"parent_id" example:"1"`
	SectionID   int64  `json:"section_id" binding:"required" example:"1"`
	AssigneeID  *int64 `json:"assignee_id" example:"1"`
	ReviewerID  *int64 `json:"reviewer_id" example:"2"`
	Title       string `json:"title" binding:"required" example:"Implement login feature"`
	Description string `json:"description" example:"Add OAuth2 login functionality"`
	Priority    int    `json:"priority" binding:"required" example:"1"`
	Deadline    *int64 `json:"deadline" example:"1737363600"`
	Status      *int   `json:"status" example:"1"`
}

type UpdateTaskRequest struct {
	ParentID    *int64 `json:"parent_id" example:"1"`
	SectionID   int64  `json:"section_id" binding:"required" example:"1"`
	AssigneeID  *int64 `json:"assignee_id" example:"1"`
	ReviewerID  *int64 `json:"reviewer_id" example:"2"`
	Title       string `json:"title" binding:"required" example:"Implement login feature"`
	Description string `json:"description" example:"Add OAuth2 login functionality"`
	Priority    int    `json:"priority" binding:"required" example:"1"`
	Deadline    *int64 `json:"deadline" example:"1737363600"`
	Status      int    `json:"status" binding:"required" example:"1"`
}

type TaskResponse struct {
	ID          int64           `json:"id" example:"1"`
	ParentID    *int64          `json:"parent_id,omitempty" example:"1"`
	SectionID   int64           `json:"section_id" example:"1"`
	Title       string          `json:"title" example:"Implement login feature"`
	Description string          `json:"description" example:"Add OAuth2 login functionality"`
	Priority    int             `json:"priority" example:"1"`
	Status      int             `json:"status" example:"1"`
	Deadline    *int64          `json:"deadline,omitempty" example:"1737363600"`
	AssigneeID  *int64          `json:"assignee_id,omitempty" example:"1"`
	Assignee    *UserResponse   `json:"assignee,omitempty"`
	ReviewerID  *int64          `json:"reviewer_id,omitempty" example:"2"`
	Reviewer    *UserResponse   `json:"reviewer,omitempty"`
	CreatedAt   int64           `json:"created_at" example:"1737277200"`
	UpdatedAt   int64           `json:"updated_at" example:"1737277200"`
	SubTasks    []*TaskResponse `json:"sub_tasks,omitempty"`
}

type TaskListResponse struct {
	Data       []*TaskResponse `json:"data"`
	TotalCount int64           `json:"total_count" example:"100"`
	Page       int             `json:"page" example:"1"`
	PageSize   int             `json:"page_size" example:"20"`
	TotalPages int             `json:"total_pages" example:"5"`
}

type CalendarEventResponse struct {
	ID           int64  `json:"id" example:"1"`
	Title        string `json:"title" example:"Backend API yaratish"`
	Description  string `json:"description" example:"REST API endpoints yaratish"`
	AssigneeID   *int64 `json:"assignee_id" example:"5"`
	AssigneeName string `json:"assignee_name" example:"Botir Nematilloyev"`
	Status       int    `json:"status" example:"2"`
	StatusName   string `json:"status_name" example:"Jarayonda"`
	StatusColor  string `json:"status_color" example:"#f59e0b"`
	StartDate    int64  `json:"start_date" example:"1704067200"`
	EndDate      *int64 `json:"end_date" example:"1704153600"`
	Deadline     *int64 `json:"deadline" example:"1704326400"`
	Priority     int    `json:"priority" example:"3"`
	SectionID    int64  `json:"section_id" example:"1"`
	SectionName  string `json:"section_name" example:"IT"`
}

func ToTaskResponse(task *domain.Task) *TaskResponse {
	resp := &TaskResponse{
		ID:          task.ID,
		SectionID:   task.SectionID,
		Title:       task.Title,
		Description: task.Description,
		Priority:    task.Priority,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
	if task.ParentID.Valid {
		resp.ParentID = &task.ParentID.Int64
	}
	if task.AssigneeID.Valid {
		resp.AssigneeID = &task.AssigneeID.Int64
	}
	if task.ReviewerID.Valid {
		resp.ReviewerID = &task.ReviewerID.Int64
	}
	if task.Deadline.Valid {
		resp.Deadline = &task.Deadline.Int64
	}
	if task.Assignee != nil {
		resp.Assignee = ToUserResponse(task.Assignee)
	}
	if task.Reviewer != nil {
		resp.Reviewer = ToUserResponse(task.Reviewer)
	}
	return resp
}

func ToTaskResponseList(tasks []*domain.Task) []*TaskResponse {
	responses := make([]*TaskResponse, len(tasks))
	for i, task := range tasks {
		responses[i] = ToTaskResponse(task)
	}
	return responses
}

func ToCalendarEventResponse(task *domain.Task, status *taskstatus.TaskStatus, sec *section.Section) *CalendarEventResponse {
	resp := &CalendarEventResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		StartDate:   task.CreatedAt, // Assuming CreatedAt as start_date if no specific field
		Priority:    task.Priority,
		SectionID:   task.SectionID,
	}

	if task.AssigneeID.Valid {
		resp.AssigneeID = &task.AssigneeID.Int64
	}
	if task.Assignee != nil {
		resp.AssigneeName = task.Assignee.FullName
	}
	if status != nil {
		resp.StatusName = status.Name
		resp.StatusColor = status.Color
	}
	if sec != nil {
		resp.SectionName = sec.Name
	}
	if task.Deadline.Valid {
		resp.Deadline = &task.Deadline.Int64
	}

	// For end_date, if it's finished, we use the value from task.EndDate which was populated from task_histories
	if task.EndDate != nil {
		resp.EndDate = task.EndDate
	} else if status != nil && status.StatusType == taskstatus.Finished {
		resp.EndDate = &task.UpdatedAt
	}

	return resp
}

func ToCalendarEventResponseList(tasks []*domain.Task, statuses map[int64]*taskstatus.TaskStatus, sections map[int64]*section.Section) []*CalendarEventResponse {
	responses := make([]*CalendarEventResponse, len(tasks))
	for i, task := range tasks {
		responses[i] = ToCalendarEventResponse(task, statuses[int64(task.Status)], sections[task.SectionID])
	}
	return responses
}

func ToTaskListResponse(result *domain.TaskListResult, page, pageSize int) *TaskListResponse {
	totalPages := int(result.TotalCount) / pageSize
	if int(result.TotalCount)%pageSize > 0 {
		totalPages++
	}

	return &TaskListResponse{
		Data:       ToTaskResponseList(result.Tasks),
		TotalCount: result.TotalCount,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}
}

// ToHierarchicalTaskListResponse converts tasks to hierarchical structure with subtasks nested in parent tasks
func ToHierarchicalTaskListResponse(result *domain.TaskListResult, page, pageSize int) *TaskListResponse {
	totalPages := int(result.TotalCount) / pageSize
	if int(result.TotalCount)%pageSize > 0 {
		totalPages++
	}

	// Build task map and identify parents and children
	taskMap := make(map[int64]*TaskResponse)
	var parentTasks []*TaskResponse

	// First pass: convert all tasks to TaskResponse and build map
	for _, task := range result.Tasks {
		taskResp := ToTaskResponse(task)
		taskResp.SubTasks = []*TaskResponse{} // Initialize empty subtasks array
		taskMap[task.ID] = taskResp
	}

	// Second pass: organize hierarchy
	for _, task := range result.Tasks {
		taskResp := taskMap[task.ID]
		if task.ParentID.Valid {
			// This is a subtask, add it to parent's SubTasks
			if parent, exists := taskMap[task.ParentID.Int64]; exists {
				parent.SubTasks = append(parent.SubTasks, taskResp)
			} else {
				// If parent is not in the current page, we still want to show this task at root or handle it?
				// Usually, in hierarchical view, if we filter by section, all tasks are returned.
				// If it's a subtask but parent is not found in the current set, treat it as a root for this view
				parentTasks = append(parentTasks, taskResp)
			}
		} else {
			// This is a parent task
			parentTasks = append(parentTasks, taskResp)
		}
	}

	return &TaskListResponse{
		Data:       parentTasks,
		TotalCount: result.TotalCount,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}
}
