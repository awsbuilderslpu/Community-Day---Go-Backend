# SCD Event Management System - Backend

## 📌 Overview
The **SCD (Student Community Day) Event Management System** is a comprehensive backend solution designed to handle end-to-end operations for large-scale events. From seamless attendee registration and role-based access control to pass generation, QR check-ins, and dynamic website content management, this system provides a robust and scalable architecture for event organizers.

## ✨ Features & Modules

### 1. 🔐 Authentication & Role-Based Access Control (RBAC)
- Secure authentication using **JWT**.
- Granular role-based permissions (Super Admin, Admin, Volunteer).
- Volunteer account management (Enable/Disable).
- Login activity tracking and audit trails.

### 2. 📊 Centralized Dashboard
- Live event statistics and real-time attendee counts.
- Track registrations (Total, Paid, Pending/Failed).
- Monitor pass distribution and check-ins.
- Recent activities and quick action shortcuts.

### 3. 🎟️ Registration & Payment
- End-to-end registration management.
- Payment status tracking and manual verification.
- Search, filter, and view detailed registration histories.

### 4. 🧑‍🤝‍🧑 Attendee Management
- Complete profile view and editing capabilities.
- Track issued passes, payment details, and check-in statuses.
- Attendee timeline and interaction history.

### 5. 🖨️ Pass Generation (Batch Processing)
- Generate passes in bulk (6 passes per A4 sheet, Landscape).
- Unique encrypted QR code for every pass.
- Store generated UIDs with batch history.
- Download previously generated PDFs.

### 6. 🎫 Pass Issuance & Tracking
- Link pre-printed passes with attendees via UID.
- Mark attendee kits as issued.
- Pass lifecycle management: Reissue, Revoke, Blacklist.
- Detailed pass issuance logs.

### 7. 📲 QR Check-In System
- Fast, secure entry validation via QR scanning.
- Duplicate entry and revoked/blacklisted pass detection.
- Instant success/error responses with entry timestamps.

### 8. 🤝 Volunteer Management
- Create and manage volunteer accounts.
- Assign specific operational roles and responsibilities.
- Track individual volunteer activities and performance.

### 9. 🔔 Notifications & Announcements
- Integrated with **Firebase Cloud Messaging (FCM)**.
- Broadcast announcements to attendees, volunteers, or admins.
- Targeted messaging and notification history.

### 10. 📅 Agenda & Schedule Management
- Dynamically manage event sessions, venues, and speakers.
- Real-time updates reflected on the event website.

### 11. 🎤 Speaker Management
- Complete speaker profiles (Bio, Image, Designation, Social Links).
- Assign speakers to specific agenda sessions.

### 12. 🌐 Dynamic Website Content Management
- Update website content without redeployment.
- Manage Hero sections, About, Venue info, FAQs, Sponsors, and Partners.

### 13. 📋 Comprehensive Audit Logs
- Track all critical actions for security and accountability.
- Log volunteer logins, pass issuance/revocation, check-ins, and content modifications.

### 14. 📈 Reports & Data Export
- Downloadable insights and reports (CSV/Excel).
- Registration, Payment, Pass Issuance, Check-in, and Volunteer Activity reports.

## 🛠️ Technology Stack
*Note: Update this section based on your finalized backend stack.*
- **Language/Framework:** Go (Golang)
- **Database:** PostgreSQL / MongoDB
- **Authentication:** JWT
- **Push Notifications:** Firebase Cloud Messaging (FCM)
- **Deployment:** AWS / Serverless

## 🚀 Getting Started

### Prerequisites
- [Go 1.21+](https://go.dev/doc/install)
- Database (PostgreSQL or similar)
- Firebase Admin SDK credentials

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/your-org/scd-event-management-backend.git
   cd scd-event-management-backend
   ```
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Set up environment variables:
   ```bash
   cp .env.example .env
   # Configure your database, JWT secret, and Firebase credentials
   ```
4. Start the development server:
   ```bash
   go run main.go
   ```

## 🤝 Contributing
1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License
This project is licensed under the MIT License - see the LICENSE file for details.
