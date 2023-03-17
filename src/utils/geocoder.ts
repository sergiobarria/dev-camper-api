import NodeGeocoder from 'node-geocoder'
import config from 'config'

const PROVIDER = config.get<string>('GEOCODER_PROVIDER')
const API_KEY = config.get<string>('GEOCODER_API_KEY')

const options = {
    provider: PROVIDER,
    httpAdapter: 'https',
    apiKey: API_KEY,
    formatter: null
}

export const geocoder = NodeGeocoder(options as any)
