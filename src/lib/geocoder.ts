import NodeGeocoder, { type Options } from 'node-geocoder'
import config from 'config'

const provider = config.get<string>('GEOCODER_PROVIDER')
const apiKey = config.get<string>('GEOCODER_CONSUMER_KEY')

const options: Options = {
    provider: provider as any,
    apiKey,
    formatter: null,
}

export const geocoder = NodeGeocoder(options)
