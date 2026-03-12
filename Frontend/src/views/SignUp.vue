<template>
    <div class="signup-container">
      <Card class="signup-card">
        <template #title>
          <h2>Sign Up</h2>
        </template>
        <template #content>
          <form @submit.prevent="handleSubmit">
            <div class="form-group">
              <label for="firstName">First Name</label>
              <InputText 
                id="firstName" 
                v-model="form.firstName" 
                required
                placeholder="Enter your first name" 
              />
              <span v-if="errors.firstName" class="error">{{ errors.firstName }}</span>
            </div>
  
            <div class="form-group">
              <label for="lastName">Last Name</label>
              <InputText 
                id="lastName" 
                v-model="form.lastName" 
                required
                placeholder="Enter your last name" 
              />
              <span v-if="errors.lastName" class="error">{{ errors.lastName }}</span>
            </div>
  
            <!-- Email input -->
            <div class="form-group">
              <label for="email">Email</label>
              <InputText 
                id="email" 
                v-model="form.email" 
                required
                placeholder="Enter your email" 
              />
              <span v-if="errors.email" class="error">{{ errors.email }}</span>
            </div>
  
            <!-- Phone number input -->
            <div class="form-group">
              <label for="phone">Phone Number</label>
              <InputText 
                id="phone" 
                v-model="form.phone" 
                required
                placeholder="Enter your phone number" 
              />
              <span v-if="errors.phone" class="error">{{ errors.phone }}</span>
            </div>
  
            <!-- Gender input -->
            <div class="form-group">
              <label for="gender">Gender</label>
              <Dropdown 
                id="gender" 
                v-model="form.gender" 
                :options="genderOptions" 
                placeholder="Select Gender" 
                required 
              />
              <span v-if="errors.gender" class="error">{{ errors.gender }}</span>
            </div>
  
            <div class="form-group">
              <label for="role">Role</label>
              <Dropdown 
                id="role" 
                v-model="form.role" 
                :options="roleOptions" 
                placeholder="Select role" 
                required 
              />
              <span v-if="errors.role" class="error">{{ errors.role }}</span>
            </div>
  
            <!-- Date of Birth input -->
            <div class="form-group">
              <label for="dob">Date of Birth</label>
              <DatePicker 
                id="dob" 
                v-model="form.dob" 
                required 
                placeholder="Select Date" 
              />
              <span v-if="errors.dob" class="error">{{ errors.dob }}</span>
            </div>
  
            <!-- Password input -->
            <div class="form-group">
              <label for="password">Password</label>
              <Password 
                id="password" 
                v-model="form.password" 
                required 
                toggleMask 
                placeholder="Enter your password" 
              />
              <span v-if="errors.password" class="error">{{ errors.password }}</span>
            </div>
  
            <Button type="submit" label="Sign Up" class="p-button-block" />
          </form>
        </template>
      </Card>
    </div>
  </template>
  
  
  <script>
import { ref } from 'vue';
import InputText from 'primevue/inputtext';
import Password from 'primevue/password';
import Button from "primevue/button";
import Card from "primevue/card";
import Dropdown from 'primevue/dropdown';
import DatePicker from 'primevue/datepicker';
import { apiCallpost } from '../plugins/testAxios';
import { v4 as uuidv4 } from 'uuid';

export default {
  components: {
    InputText,
    Password,
    Button,
    Card,
    Dropdown,
    DatePicker
  },
  data() {
    return {
      form: {
        firstName: '',
        lastName: '',
        email: '',
        phone: '',
        gender: '',
        role: '',
        dob: '',
        password: '',
      },
      genderOptions: [
        "male", "female", "other"
      ],
      roleOptions: [
        "customer", "Admin"
      ],
      errors: {}, // Error tracking
    };
  },
  methods: {
    // Form validation
    validateForm() {
      this.errors = {}; // Reset errors
      let isValid = true;

      // Validate firstName
      if (!this.form.firstName) {
        this.errors.firstName = 'First name is required.';
        isValid = false;
      }

      // Validate lastName
      if (!this.form.lastName) {
        this.errors.lastName = 'Last name is required.';
        isValid = false;
      }

      // Validate email
      const emailPattern = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
      if (!this.form.email) {
        this.errors.email = 'Email is required.';
        isValid = false;
      } else if (!emailPattern.test(this.form.email)) {
        this.errors.email = 'Please enter a valid Gmail address.';
        isValid = false;
      }

      // Validate phone
      if (!this.form.phone) {
        this.errors.phone = 'Phone number is required.';
        isValid = false;
      }

      // Validate gender
      if (!this.form.gender) {
        this.errors.gender = 'Gender is required.';
        isValid = false;
      }

      // Validate role
      if (!this.form.role) {
        this.errors.role = 'Role is required.';
        isValid = false;
      }

      // Validate dob
      if (!this.form.dob) {
        this.errors.dob = 'Date of Birth is required.';
        isValid = false;
      }

      // Validate password
      if (!this.form.password) {
        this.errors.password = 'Password is required.';
        isValid = false;
      } else if (this.form.password.length < 5) {
        this.errors.password = 'Password should be at least 5 characters.';
        isValid = false;
      }

      return isValid;
    },

    handleSubmit() {
      if (this.validateForm()) {
        this.userData = JSON.parse(JSON.stringify(this.form));
        const userId = uuidv4();
        this.userData.userId = userId;

        const payload = this.userData;
         const response = fetch("http://localhost:8000/o/register", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(payload),
      })
        .then((res) => res.json())
        .then((data) => {
          console.log("Registration successful:", data);
          alert("Registration successful! Please log in.");
          this.resetForm();
        })
        .catch((error) => {
          console.error("Error during registration:", error);
          alert("An error occurred during registration. Please try again.");
        });
      } else {
        console.error("Form validation failed:", this.errors);
      }
    },

    resetForm() {
      this.form = {
        firstName: '',
        lastName: '',
        email: '',
        phone: '',
        gender: '',
        role: '',
        dob: '',
        password: '',
      };
    },
  },
};
</script>

  
  <style scoped>
  .signup-container {
    max-width: 500px;
    margin: 0 auto;
    padding: 40px;
    background-image: url('https://via.placeholder.com/1500'); /* Set your background image here */
    background-size: cover;
    background-position: center;
    min-height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
  }
  
  .signup-card {
    padding: 30px;
    background-color: rgba(255, 255, 255, 0.85); /* Semi-transparent background for card */
    border-radius: 10px;
    box-shadow: 0 0 15px rgba(0, 0, 0, 0.1);
  }
  
  h2 {
    text-align: center;
    margin-bottom: 20px;
  }
  
  .form-group {
    margin-bottom: 15px;
  }
  
  label {
    font-weight: bold;
    margin-bottom: 5px;
    display: block;
  }
  
  input,
  select {
    width: 100%;
    padding: 10px;
    margin-bottom: 10px;
    border-radius: 5px;
    border: 1px solid #ddd;
  }
  
  button {
    width: 100%;
    padding: 10px;
    background-color: #007bff;
    color: white;
    border: none;
    border-radius: 5px;
    cursor: pointer;
  }
  
  button:hover {
    background-color: #0056b3;
  }
  </style>
  