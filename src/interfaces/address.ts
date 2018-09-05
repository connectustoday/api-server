import { prop, Typegoose, ModelType, InstanceType } from 'typegoose';
import IPoint from './point';

export default class IAddress {

	@prop({ required: true })
	schema_version: number;

	@prop()
	street?: string;

	@prop()
	city?: string;

	@prop()
	province?: string;

	@prop()
	country?: string;

	@prop()
	postal_code?: string;

	@prop()
	apt_number?: string;

	@prop()
	geojson?: IPoint;
}
