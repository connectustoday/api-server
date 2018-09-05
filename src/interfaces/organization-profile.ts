import { prop, Typegoose, ModelType, InstanceType } from 'typegoose';
import IAddress from './address';

export default class IOrganizationProfile {
	@prop({ required: true })
	schema_version: string;

	@prop()
	mission?: string;

	@prop()
	quote?: string;

	@prop()
	address?: IAddress;

	@prop()
	affiliated_orgs?: Array<string>;

	@prop()
	interests?: Array<string>;
}
