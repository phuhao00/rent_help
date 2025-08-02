import axios from 'axios';
import { TokenResponse, LoginRequest, RegisterRequest, User, Property, Booking } from '@/types';

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8081/api/v1';

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Add auth interceptor
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('accessToken');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

// Auth API
export const authApi = {
  login: (data: LoginRequest): Promise<TokenResponse> =>
    api.post('/auth/login', data).then(res => res.data),
  
  register: (data: RegisterRequest): Promise<{ message: string; user: User }> =>
    api.post('/auth/register', data).then(res => res.data),
  
  refreshToken: (): Promise<TokenResponse> =>
    api.post('/auth/refresh').then(res => res.data),
};

// User API
export const userApi = {
  getProfile: (): Promise<User> =>
    api.get('/users/profile').then(res => res.data),
  
  updateProfile: (data: Partial<User>): Promise<{ message: string }> =>
    api.put('/users/profile', data).then(res => res.data),
};

// Property API
export const propertyApi = {
  getProperties: (params?: {
    limit?: number;
    skip?: number;
    city?: string;
    type?: string;
  }): Promise<Property[]> =>
    api.get('/properties', { params }).then(res => res.data),
  
  getProperty: (id: string): Promise<Property> =>
    api.get(`/properties/${id}`).then(res => res.data),
  
  createProperty: (data: Omit<Property, 'id' | 'owner_id' | 'created_at' | 'updated_at'>): Promise<Property> =>
    api.post('/properties', data).then(res => res.data),
  
  updateProperty: (id: string, data: Partial<Property>): Promise<{ message: string }> =>
    api.put(`/properties/${id}`, data).then(res => res.data),
  
  deleteProperty: (id: string): Promise<{ message: string }> =>
    api.delete(`/properties/${id}`).then(res => res.data),
};

// Booking API
export const bookingApi = {
  getBookings: (params?: {
    limit?: number;
    skip?: number;
    status?: string;
  }): Promise<Booking[]> =>
    api.get('/bookings', { params }).then(res => res.data),
  
  getBooking: (id: string): Promise<Booking> =>
    api.get(`/bookings/${id}`).then(res => res.data),
  
  createBooking: (data: Omit<Booking, 'id' | 'tenant_id' | 'status' | 'created_at' | 'updated_at'>): Promise<Booking> =>
    api.post('/bookings', data).then(res => res.data),
  
  updateBooking: (id: string, data: Partial<Booking>): Promise<{ message: string }> =>
    api.put(`/bookings/${id}`, data).then(res => res.data),
  
  deleteBooking: (id: string): Promise<{ message: string }> =>
    api.delete(`/bookings/${id}`).then(res => res.data),
};

export default api;
