package database

const (
	createSubscription = "create subscription"
	deleteSubscription = "delete subscription by uuid"
	getSubscription    = "get subscription by uuid"
	listSubscription   = "list subscriptions"
	updateSubscription = "update subscription by uuid"
)

func queriesSubscription() map[string]string {
	return map[string]string{
		createSubscription: `INSERT INTO subscriptions (user_id, classroom_id, role, expires_at) 
			VALUES (:user_id, :classroom_id, :role, :expires_at) 
			RETURNING *`,
		deleteSubscription: `UPDATE subscriptions SET deleted_at = NOW() WHERE uuid = :uuid`,
		getSubscription:    "SELECT * FROM subscriptions WHERE uuid = :uuid",
		listSubscription:   "SELECT * FROM subscriptions",
		updateSubscription: `UPDATE subscriptions 
			SET user_id = :user_id, classroom_id = :classroom_id, expires_at = :expires_at 
			WHERE id = :uuid 
			RETURNING *`,
	}
}
