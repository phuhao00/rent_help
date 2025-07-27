export interface User {
  id: string;
  email: string;
  name: string;
  phone?: string;
  avatar?: string;
  role: 'tenant' | 'landlord' | 'admin';
  verified: boolean;
  created_at: string;
  updated_at: string;
}

export interface Property {
  id: string;
  title: string;
  description: string;
  type: 'apartment' | 'house' | 'room';
  price: number;
  currency: string;
  location: Location;
  features: string[];
  images: string[];
  available: boolean;
  owner_id: string;
  created_at: string;
  updated_at: string;
}

export interface Location {
  address: string;
  city: string;
  state: string;
  country: string;
  latitude: number;
  longitude: number;
}

export interface Booking {
  id: string;
  property_id: string;
  tenant_id: string;
  start_date: string;
  end_date: string;
  status: 'pending' | 'confirmed' | 'cancelled' | 'completed';
  total_price: number;
  message?: string;
  created_at: string;
  updated_at: string;
}

export interface LoginRequest {
  email: string;
  password: string;
}

export interface RegisterRequest {
  email: string;
  password: string;
  name: string;
  phone?: string;
  role: 'tenant' | 'landlord';
}

export interface TokenResponse {
  access_token: string;
  refresh_token?: string;
  expires_in: number;
}
