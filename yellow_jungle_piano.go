package main

import (
	"fmt"
	"net/http"
)

// Server struct
type Server struct {
	mentors []Mentor
	network []Student
}

// Mentor Struct for mentor
type Mentor struct {
	name string
	school string
	email string
	newsletter bool
}

// Student struct
type Student struct {
	name string
	school string
	email string
}

// RegisterMentor : Handler for registering mentors into the system
func (s *Server) RegisterMentor(w http.ResponseWriter, r *http.Request) {

	mentor := Mentor{
		name: r.FormValue("name"),
		school: r.FormValue("school"),
		email: r.FormValue("email"),
		newsletter: r.FormValue("newsletter") == "true",
	}

	s.mentors = append(s.mentors, mentor)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Mentor %s registered successfully", mentor.name)
}

// RegisterStudent : Handler for registering students into the system
func (s *Server) RegisterStudent(w http.ResponseWriter, r *http.Request) {
	
	student := Student{
		name: r.FormValue("name"),
		school: r.FormValue("school"),
		email: r.FormValue("email"),
	}

	s.network = append(s.network, student)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Student %s registered successfully", student.name)
}

// GetMentor : Handler for retrieving a single mentor
func (s *Server) GetMentor(w http.ResponseWriter, r *http.Request, mentorID string) {

	for _, mentor := range s.mentors {
		if mentor.email == mentorID {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Mentor: %s, School: %s, Email: %s, Newsletter: %t", mentor.name, mentor.school, mentor.email, mentor.newsletter)
			return
		}
	}
	
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Mentor with email %s not found", mentorID)
}

// GetMentorList : Handler for retrieving a list of mentors
func (s *Server) GetMentorList(w http.ResponseWriter, r *http.Request) {
	
	if len(s.mentors) == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "No mentors have been registered")
		return
	}

	for _, mentor := range s.mentors {
		fmt.Fprintf(w, "Mentor: %s, School: %s, Email: %s, Newsletter: %t\n", mentor.name, mentor.school, mentor.email, mentor.newsletter)
	}
}

// GetStudentList : Handler for retrieving a list of students
func (s *Server) GetStudentList(w http.ResponseWriter, r *http.Request) {
	
	if len(s.network) == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "No students have been registered")
		return
	}

	for _, student := range s.network {
		fmt.Fprintf(w, "Student: %s, School: %s, Email: %s\n", student.name, student.school, student.email)
	}
}

// AssignMentor : Handler for assigning a mentor to a student
func (s *Server) AssignMentor(w http.ResponseWriter, r *http.Request) {
	
	studentID := r.FormValue("studentID")
	mentorID := r.FormValue("mentorID")

	var student *Student
	var mentor *Mentor

	// Find student
	for i, networkStudent := range s.network {
		if networkStudent.email == studentID {
			student = &s.network[i]
			break
		}
	}

	// Find mentor
	for i, mentorStudent := range s.mentors {
		if mentorStudent.email == mentorID {
			mentor = &s.mentors[i]
			break
		}
	}

	if student == nil || mentor == nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Student or mentor with specified ID not found")
		return
	}

	student.name = mentor.name
	student.school = mentor.school
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Mentor %s has been assigned to student %s", mentor.name, student.name)
}

// SendNewsletter : Handler for sending the newsletter to subscribed mentors
func (s *Server) SendNewsletter(w http.ResponseWriter, r *http.Request) {
	
	for _, mentor := range s.mentors {
		if mentor.newsletter {
			// code to send the newsletter to mentor
			fmt.Fprintf(w, "Newsletter sent to mentor %s", mentor.name)
		}
	}
	
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Newsletter sent to all subscribed mentors")
}

// InitServer : Initializes the server
func InitServer() *Server {
	return &Server{}
}

func main() {
	server := InitServer()

	http.HandleFunc("/register/mentor", server.RegisterMentor)
	http.HandleFunc("/register/student", server.RegisterStudent)
	http.HandleFunc("/mentors/", server.GetMentor)
	http.HandleFunc("/mentors/list", server.GetMentorList)
	http.HandleFunc("/students/list", server.GetStudentList)
	http.HandleFunc("/mentors/assign", server.AssignMentor)
	http.HandleFunc("/mentors/newsletter", server.SendNewsletter)
	
	http.ListenAndServe(":8080", nil)
}