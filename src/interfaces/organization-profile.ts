import { prop, Typegoose, ModelType, InstanceType } from 'typegoose';

class IOrganizationProfile {
	@prop({ required: true })
	schema_version: string;

	@prop()
	mission?: string;

	@prop()
	quote?: string;

	@prop()
	address?: string; //TODO

	@prop()
	affiliated_orgs?: Array<string>;

	@prop()
	interests?: Array<string>;
}
