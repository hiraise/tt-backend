package persistent

import (
	"context"
	"fmt"
	"strings"
	"task-trail/internal/usecase/dto"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PgProjectRepository struct {
	PgRepostitory
}

func NewProjectRepo(db *pgxpool.Pool) *PgProjectRepository {
	return &PgProjectRepository{PgRepostitory{pg: db}}
}

func (r *PgProjectRepository) Create(ctx context.Context, data *dto.ProjectCreate) (int, error) {
	query := `
		INSERT INTO projects
		(owner_id, name, description)
		VALUES ($1, $2, $3)
		RETURNING id;`
	var id int
	err := r.getDb(ctx).QueryRow(ctx, query, data.OwnerID, data.Name, data.Description).Scan(&id)
	if err != nil {
		return 0, r.handleError(err)
	}
	return id, nil
}

func (r *PgProjectRepository) GetList(ctx context.Context, data *dto.ProjectList) ([]*dto.ProjectRes, error) {
	query := `
		SELECT P.id, P.name, P.description, COUNT(T.id)
		FROM public.projects as P 
		LEFT JOIN public.tasks as T on P.id = T.project_id
		WHERE P.id IN (SELECT project_id FROM project_users WHERE user_id = $1) AND 
		GROUP BY (P.id)
	`
	rows, err := r.getDb(ctx).Query(ctx, query, data.MemberID)
	if err != nil {
		return nil, r.handleError(err)
	}
	defer rows.Close()

	var retVal []*dto.ProjectRes
	for rows.Next() {
		var item dto.ProjectRes
		if err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.TaskCount); err != nil {
			return nil, r.handleError(err)
		}
		retVal = append(retVal, &item)
	}
	if err := rows.Err(); err != nil {
		return nil, r.handleError(err)
	}
	return retVal, nil
}

func (r *PgProjectRepository) AddMembers(ctx context.Context, data *dto.ProjectAddMembersDB) error {
	values := make([]any, 0, len(data.MemberIDs)+1)
	values = append(values, data.ProjectID)
	var items []string
	for i, id := range data.MemberIDs {
		values = append(values, id)
		items = append(items, fmt.Sprintf("($%d, $1)", i+2))
	}
	query := fmt.Sprintf(`
		INSERT INTO project_users 
		(user_id, project_id)
		VALUES
		%s
		`, strings.Join(items, ",\n"))
	_, err := r.getDb(ctx).Exec(ctx, query, values...)
	if err != nil {
		return r.handleError(err)
	}
	return nil
}

func (r *PgProjectRepository) GetOwnedProject(ctx context.Context, projectID int, ownerID int) (*dto.Project, error) {
	query := `
		SELECT id, name, description, owner_id, created_at FROM projects WHERE id = $1 AND owner_id = $2;
	`
	var item dto.Project
	err := r.getDb(ctx).QueryRow(ctx, query, projectID, ownerID).Scan(&item.ID, &item.Name, &item.Description, &item.OwnerID, &item.CreatedAt)
	if err != nil {
		return nil, r.handleError(err)
	}
	query = `
		SELECT id, email FROM users WHERE id IN (SELECT user_id FROM project_users WHERE project_id = $1);
		`
	rows, err := r.getDb(ctx).Query(ctx, query, projectID)
	if err != nil {
		return nil, r.handleError(err)
	}

	defer rows.Close()

	for rows.Next() {
		var u dto.UserEmailAndID
		if err := rows.Scan(&u.ID, &u.Email); err != nil {
			return nil, r.handleError(err)
		}
		item.Members = append(item.Members, &u)

	}
	if err := rows.Err(); err != nil {
		return nil, r.handleError(err)
	}

	return &item, nil

}
