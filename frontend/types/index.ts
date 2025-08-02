export interface User {
  id: string;
  email: string;
  full_name: string;
  first_name?: string;
  last_name?: string;
  phone?: string;
  avatar?: string;
  role: 'tenant' | 'landlord' | 'admin';
  is_verified: boolean;
  verified: boolean;
  is_active: boolean;
  created_at: string;
  updated_at: string;
}

export interface Property {
  id: string;
  title: string;
  description: string;
  type: 'apartment' | 'house' | 'condo' | 'townhouse' | 'studio';
  price: number;
  currency?: string;
  address: Address;
  location: GeoLocation;
  bedrooms: number;
  bathrooms: number;
  area: number;
  amenities: string[];
  images: string[];
  available: boolean;
  available_from: string;
  lease_terms: string[];
  pets_allowed: boolean;
  smoking_allowed: boolean;
  utilities_included: string[];
  owner_id: string;
  created_at: string;
  updated_at: string;
}

export interface Address {
  street: string;
  city: string;
  state: string;
  zip_code: string;
  country: string;
  latitude?: number;
  longitude?: number;
}

export interface GeoLocation {
  type: string;
  coordinates: number[];
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
  user?: User;
}
