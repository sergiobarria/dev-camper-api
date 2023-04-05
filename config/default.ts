import dotenv from 'dotenv'

dotenv.config()

export default {
    PORT: process.env.PORT,
    NODE_ENV: process.env.NODE_ENV,
    GEOCODER_PROVIDER: process.env.GEOCODER_PROVIDER,
    GEOCODER_CONSUMER_KEY: process.env.GEOCODER_CONSUMER_KEY,
}
