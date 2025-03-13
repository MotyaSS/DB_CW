package service

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/MotyaSS/DB_CW/pkg/config"
)

type BackupService struct {
	DBName     string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	BackupDir  string
}

func NewBackupService(config config.Database) *BackupService {
	return &BackupService{
		DBName:     config.DBName,
		DBUser:     config.Username,
		DBPassword: config.Password,
		DBHost:     config.Host,
		DBPort:     config.Port,
		BackupDir:  config.BackupPath,
	}
}

func (s *BackupService) CreateBackup() (string, error) {
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	filename := fmt.Sprintf("%s_%s.sql", s.DBName, timestamp)
	backupPath := filepath.Join(s.BackupDir, filename)
	fmt.Println(backupPath)
	if err := os.MkdirAll(s.BackupDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create backup directory: %w", err)
	}

	cmd := exec.Command("pg_dump",
		"-h", s.DBHost,
		"-p", s.DBPort,
		"-U", s.DBUser,
		"-F", "p",
		"-f", backupPath,
		s.DBName,
	)

	cmd.Env = append(os.Environ(), fmt.Sprintf("PGPASSWORD=%s", s.DBPassword))

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to create backup: %w", err)
	}

	return backupPath, nil
}

func (s *BackupService) RestoreFromBackup(backupName string) error {
	backupPath := filepath.Join(s.BackupDir, backupName)
	if _, err := os.Stat(backupPath); os.IsNotExist(err) {
		return fmt.Errorf("backup file does not exist: %s", backupName)
	}

	cmd := exec.Command("psql",
		"-h", s.DBHost,
		"-p", s.DBPort,
		"-U", s.DBUser,
		"-d", s.DBName,
		"-f", backupPath,
	)

	cmd.Env = append(os.Environ(), fmt.Sprintf("PGPASSWORD=%s", s.DBPassword))

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to restore backup: %w", err)
	}

	return nil
}

func (s *BackupService) ListBackups() ([]string, error) {
	files, err := os.ReadDir(s.BackupDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read backup directory: %w", err)
	}

	backups := []string{}
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".sql" {
			backups = append(backups, file.Name())
		}
	}

	return backups, nil
}
